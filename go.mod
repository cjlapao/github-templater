module github.com/cjlapao/github-templater

go 1.21

require (
	github.com/cjlapao/common-go v0.0.39
	github.com/cjlapao/common-go-dependency-tree v0.2.0
	github.com/cjlapao/common-go-logger v0.0.5
	github.com/pkg/errors v0.9.1

)

require (
	// indirect
	github.com/fatih/color v1.16.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.20 // indirect
	golang.org/x/sys v0.17.0 // indirect
	gopkg.in/yaml.v3 v3.0.1
)

// replace github.com/cjlapao/common-go-dependency-tree => ../common_framework/common-go-dependency-tree
