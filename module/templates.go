package module

const fileTpl = `// Code generated by protoc-gen-rtk-query. DO NOT EDIT.
// source: {{ .file.InputPath }}
{{- $fn := fname .file.Name}}
{{- $fnl := fnamelc .file.Name}}

import { createApi } from '@reduxjs/toolkit/query/react'
import { GrpcWebFetchTransport, GrpcWebOptions } from '@protobuf-ts/grpcweb-transport';
{{- if withMetadata }}
import { grpcBaseQuery, providesList, GrpcBaseQueryMeta, WithMetadata } from '@greatliontech/protobuf-ts-rtk-query';
{{- else }}
import { grpcBaseQuery, providesList } from '@greatliontech/protobuf-ts-rtk-query';
{{- end }}
{{- range .imports}}
{{.}}
{{- end}}

export const grpcWebOptions: GrpcWebOptions = {
  baseUrl: 'http://localhost:8080'
}

const transport = new GrpcWebFetchTransport(grpcWebOptions);

{{- range .file.Services }}
{{- $sn := .Name }}
{{- $snl := .Name.LowerCamelCase }}

const {{$snl}}Client = new {{$sn}}Client(transport)

// Define a service using a base URL and expected endpoints
export const {{$snl}} = createApi({
  reducerPath: '{{$snl}}',
  baseQuery: grpcBaseQuery(),
{{- if hasTags . }}
	tagTypes: {{ tags . }},
{{- end }}
  endpoints: (builder) => ({
{{- range .Methods }}
{{- $mn := .Name }}
{{- $mnl := .Name.String | lowerFirst }}
    {{$mnl}}: builder.{{endpoint .}}<{{if withMetadata}}WithMetadata<{{end}}{{importName .Output $.names}}{{if withMetadata}}>{{end}}, {{importName .Input $.names}}>({
      query: (req) => {{$snl}}Client.{{$mnl}}(req),
{{- if withMetadata }}
      transformResponse(baseQueryReturnValue: {{importName .Output $.names}}, meta: GrpcBaseQueryMeta, arg) {
        return {
          ...baseQueryReturnValue,
          grpcQueryMetadata: meta
        }
      },
{{- end }}
{{- if hasProvidesTags . }}
	    providesTags: {{ providesTags . }},
{{- end }}
{{- if hasInvalidatesTags . }}
	    invalidatesTags: {{ invalidatesTags . }},
{{- end }}
    }),
{{- end }}
  }),
})

// Export hooks for usage in function components, which are
// auto-generated based on the defined endpoints
export const {
{{- range .Methods }}
  use{{.Name}}{{ endpoint . | upperFirst }},
{{- end }}
} = {{$snl}}
{{ end }}
`
const storeTpl = `// Code generated by protoc-gen-rtk-query. DO NOT EDIT.
{{ range .Imports }}
{{.}}
{{- end }}

export const apiReducers = {
{{- range .Reducers }}
  {{.}}
{{- end }}
}

export const apiMiddlewares = [
{{- range .Middlewares }}
  {{.}}
{{- end }}
]
`
