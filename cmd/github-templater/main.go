package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/common-go/version"
	"github.com/cjlapao/github-templater/pkg/provision"
	golangcilinter "github.com/cjlapao/github-templater/pkg/provisioners/golangci_linter"
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
	p := provision.New()
	p.Add(golangcilinter.GolangCILinterProvisioner{})
	p.Add(makefile.MakeFileProvisioner{})
	p.Provision()
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
