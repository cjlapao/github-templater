package helpers

import (
	"html/template"
	"strings"

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
		return string(b)
	},
	"nindent": func(n int, s string) string {
		lines := strings.Split(s, "\n")
		for i, line := range lines {
			lines[i] = Indent(n, line)
		}
		return strings.Join(lines, "\n")
	},
}
