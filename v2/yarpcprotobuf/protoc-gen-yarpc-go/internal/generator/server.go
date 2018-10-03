package generator

const _serverTemplate = `
{{define "server" -}}
{{with .File}}

{{range .Services -}}
  {{$svc := .Name}}

  {{/* Service interface */}}
  type {{$svc}}Server interface {
    {{range .Methods -}}
      {{if and (not .ClientStreaming) (not .ServerStreaming) -}}
        {{.Name}}(context.Context, {{.Request.Name}}) ({{.Response.Name}}, error)
      {{else -}}
        {{.Name}}(
          {{if not .ClientStreaming}}
            {{.Request.Name}},
          {{end -}}
          {{$svc}}{{.Name}}ServerStream,
        ) {{if not .ServerStreaming -}} ({{$svc}}{{.Name}}ServerStream, error) {{else -}} error {{end -}}
      {{end -}}
    {{end -}}
  }


  {{/* Stream server interfaces */}}

  {{range .Methods -}}
    {{if or .ClientStreaming .ServerStreaming}}
    type {{$svc}}{{.Name}}ServerStream interface {
      Context() context.Context
    {{if .ClientStreaming -}}
      Recv(...yarpc.StreamOption) ({{.Request.Name}}, error)
    {{end -}}
    {{if .ServerStreaming -}}
      Send({{.Response.Name}}, ...yarpc.StreamOption) error
    {{end -}}
    }
    {{end -}}
  {{end -}}
{{end -}}

{{end -}}{{end -}}
`
