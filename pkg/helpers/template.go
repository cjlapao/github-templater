package helpers

import (
	"html/template"
	"strings"

	"github.com/cjlapao/common-go/helper"
	"github.com/cjlapao/github-templater/pkg/diagnostics"
	"gopkg.in/yaml.v3"
)

var FuncMap = template.FuncMap{
	"toLower": strings.ToLower,
	"toUpper": strings.ToUpper,
	"toYaml": func(v interface{}) string {
		b, err := yaml.Marshal(v)
		if err != nil {
			panic(err)
		}
		result := string(b)
		result = strings.TrimSuffix(result, "\n")
		return result
	},
	"nindent": func(n int, s string) string {
		lines := strings.Split(s, "\n")
		for i, line := range lines {
			lines[i] = Indent(n, line)
		}
		return strings.Join(lines, "\n")
	},
}

func WriteTemplateToFile(filename string, templateString string, data interface{}) diagnostics.Diagnostics {
	diag := diagnostics.New()
	// Parse the template
	tmpl, err := template.New("default").Funcs(FuncMap).Parse(templateString)
	if err != nil {
		diag.AddError(err)
		return diag
	}

	// Execute the template
	var output strings.Builder
	err = tmpl.Execute(&output, data)
	if err != nil {
		diag.AddError(err)
		return diag
	}

	err = helper.WriteToFile(output.String(), filename)
	if err != nil {
		diag.AddError(err)
		return diag
	}

	return diag
}
