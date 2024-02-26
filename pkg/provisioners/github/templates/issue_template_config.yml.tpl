blank_issues_enabled: {{ .BlankIssuesEnabled }}
{{- if and .ContactLinks (gt (len .ContactLinks) 0) }}
contact_links:
{{ .ContactLinks | toYaml }}
{{- end }}

