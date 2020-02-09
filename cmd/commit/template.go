package commit

// 错误输出模板
var errorOutputTemplate = `{{color "red"}}{{ ErrorIcon }} 对不起，您的输入有误: {{.Error}}{{color "reset"}}
`

// 选项类型问题提示模板
var selectQuestionTemplate = `
{{- if .ShowHelp }}{{- color "cyan"}}{{ HelpIcon }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color "green+hb"}}{{ QuestionIcon }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }}{{ .FilterMessage }}{{color "reset"}}
{{- if .ShowAnswer}}{{color "cyan"}} {{.Answer}}{{color "reset"}}{{"\n"}}
{{- else}}
 {{- "  "}}{{- color "cyan"}}{{color "reset"}}
 {{- "\n"}}
 {{- range $ix, $choice := .PageEntries}}
   {{- if eq $ix $.SelectedIndex}}{{color "cyan+b"}}{{ SelectFocusIcon }} {{else}}{{color "default+hb"}}  {{end}}
   {{- $choice}}
   {{- color "reset"}}{{"\n"}}
 {{- end}}
{{- end}}`

// 多行输入类型问题提示模板
var multilineQuestionTemplate = `
{{- if .ShowHelp }}{{- color "cyan"}}{{ HelpIcon }} {{ .Help }}{{color "reset"}}{{"\n"}}{{end}}
{{- color "green+hb"}}{{ QuestionIcon }} {{color "reset"}}
{{- color "default+hb"}}{{ .Message }} {{color "reset"}}
{{- if .ShowAnswer}}
  {{- "\n"}}{{color "cyan"}}{{.Answer}}{{color "reset"}}
  {{- if .Answer }}{{ "\n" }}{{ end }}
{{- else }}
  {{- if .Default}}{{color "white"}}({{.Default}}) {{color "reset"}}{{end}}
  {{- color "cyan"}}[连续输入两个空行以结束提交]{{color "reset"}}
{{- end}}`

// 结果格式化模板
var answerFormatTemplate = `
{{.Type}}{{if ne .Scope ""}}({{.Scope}}){{end}}{{": "}}{{.Short}}
{{if ne .Long ""}}{{.Long}}{{"\n"}}{{end}}
{{if ne .Breaking ""}}{{"不兼容变更：\n"}}{{.Breaking}}{{"\n"}}{{end}}
{{if ne .Issue ""}}{{"修复或关闭的issue：\n"}}{{.Issue}}{{end}}
`
