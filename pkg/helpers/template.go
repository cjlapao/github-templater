package helpers

import (
	"fmt"
	"html/template"
	"strings"

	"gopkg.in/yaml.v3"
)

var FuncMap = template.FuncMap{
	"toLower": strings.ToLower,
	"toUpper": strings.ToUpper,
	"toYaml": func(v interface{}) string {
		fmt.Printf("v: %v\n", v)
		b, err := yaml.Marshal(v)
		if err != nil {
			panic(err)
		}
		fmt.Printf("b: %v\n", string(b))
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
