package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/provision"
	"github.com/cjlapao/github-templater/pkg/provisioners/github"
	dependabot "github.com/cjlapao/github-templater/pkg/provisioners/github_dependabot"
	"github.com/cjlapao/github-templater/pkg/provisioners/golangci_linter"
	"github.com/cjlapao/github-templater/pkg/provisioners/makefile"
)

func main() {
	versionSvc := SetVersion("../VERSION")
	getVersion := helper.GetFlagSwitch("version", false)
	if getVersion {
		format := helper.GetFlagValue("o", "json")
		switch strings.ToLower(format) {
		case "json":
			fmt.Println(versionSvc.PrintVersion(int(version.JSON)))
		case "yaml":
			fmt.Println(versionSvc.PrintVersion(int(version.JSON)))
		default:
			fmt.Println("Please choose a valid format, this can be either json or yaml")
		}
		os.Exit(0)
	}

	versionSvc.PrintAnsiHeader()

	defer func() {
	}()

	Init()
}

func Init() {
	ctx := context.Get()
	p := provision.New()
	p.Add(&github.GitHubProvisioner{})
	p.Add(&golangci_linter.GolangCILinterProvisioner{})
	p.Add(&makefile.MakeFileProvisioner{})
	p.Add(&dependabot.DependabotProvisioner{})
	if diag := p.Provision(); diag.HasErrors() {
		for _, err := range diag.Errors() {
			ctx.LogError(err.Error())
		}
		ctx.LogError("There was an error in the providers")
		os.Exit(1)
	}
}

func SetVersion(ver string) *version.Version {
	versionSvc := version.Get()
	versionSvc.Name = "GitHub Templater"
	versionSvc.Author = "Carlos Lapao"
	versionSvc.License = "MIT"
	strVer, err := version.FromFile(ver)
	if err == nil {
		versionSvc.Major = strVer.Major
		versionSvc.Minor = strVer.Minor
		versionSvc.Build = strVer.Build
		versionSvc.Rev = strVer.Rev
	}

	return versionSvc
}
