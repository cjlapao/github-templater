package common

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/config"
	"github.com/cjlapao/github-templater/pkg/context"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
)

func CheckIfFileExists(ctx *context.ProvisionerContext, name string, folder string, potentialFileNames []string, envVariable string) diagnostics.Diagnostics {
	diag := diagnostics.New()
	foundFiles := []string{}
	for _, fileName := range potentialFileNames {
		if helper.FileExists(path.Join(folder, fileName)) {
			foundFiles = append(foundFiles, fileName)
		}
	}

	if len(foundFiles) > 0 {
		cfg := config.Get()
		foundFilesString := strings.Join(foundFiles, ", ")
		fileStr := "file"
		if len(foundFiles) > 1 {
			fileStr = "files"
		}
		override := cfg.RequestBoolFromUser(envVariable, fmt.Sprintf("A %v config %v \"%v\" already exists, do you want to override it? [y/n]", name, fileStr, foundFilesString), false)
		if !override {
			msg := fmt.Sprintf("A %v config %v \"%v\" already exists, ignoring provisioner on request by user", name, fileStr, foundFilesString)
			ctx.LogWarn(msg)
			diag.AddWarning(msg)
			return diag
		}

		for _, fileName := range foundFiles {
			err := os.Remove(path.Join(folder, fileName))
			if err != nil {
				diag.AddError(err)
				return diag
			}
		}
	}

	return diag
}
