daysUntilStale: {{ .Template.DaysUntilStale }}
daysUntilClose: {{ .Template.DaysUntilClose }}
{{- if .Template.ExemptLabels }}
exemptLabels:
{{- range .Template.ExemptLabels }}
- {{ . }}
{{- end }}
{{- end }}
exemptProjects: {{ .Template.ExemptProjects }}
exemptAssignees: {{ .Template.ExemptAssignees }}
{{- if .Template.MarkComment }}
markComment: {{ .Template.MarkComment | toYaml }}
{{- end }}
{{- if .Template.UnmarkComment }}
unmarkComment: {{ .Template.UnmarkComment | toYaml }}
{{- end }}
{{- if .Template.CloseComment }}
closeComment: {{ .Template.CloseComment | toYaml }}
{{- end }}
limitPerRun: {{ .Template.LimitPerRun }}
{{- if .Template.Only }}
only: {{ .Template.Only }}
{{- end }}
{{- if .Pulls }}
pulls: 
{{ .Pulls | toYaml | nindent 2}}
{{- end }}
{{- if .Issues }}
issues:
{{ .Issues | toYaml | nindent 2}}
{{- end }}