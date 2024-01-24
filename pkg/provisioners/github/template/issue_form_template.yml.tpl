name: {{ .Name | toLower }}
{{- if .Description }}
description: {{ .Description }}
{{- end }}
{{- if and .Labels (gt (len .Labels) 0) }}
labels:
{{ .Labels | toYaml }}
{{- end }}
{{- if and .Projects (gt (len .Projects) 0) }}
projects:
{{ .Projects | toYaml }}
{{- end }}
{{- if and .Assignees (gt (len .Assignees) 0) }}
assignees:
{{ .Assignees | toYaml }}
{{- end }}
{{- if .Body }}
body:
{{ .Body.Items | toYaml }}
{{- end }}
