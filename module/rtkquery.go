package module

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	"github.com/lyft/protoc-gen-star"

	rtkquerypb "github.com/greatliontech/protoc-gen-rtk-query/proto/rtkquery"
)

type Module struct {
	*pgs.ModuleBase
}

type storeFile struct {
	Imports     []string
	Reducers    []string
	Middlewares []string
}

func RTKQuery() pgs.Module { return &Module{ModuleBase: &pgs.ModuleBase{}} }

func (m *Module) InitContext(ctx pgs.BuildContext) {
	m.ModuleBase.InitContext(ctx)
}

func (m *Module) Name() string { return "rtk-query" }

func (m *Module) Execute(targets map[string]pgs.File, pkgs map[string]pgs.Package) []pgs.Artifact {

	tpl := template.New("")

	sprigFuncMap := sprig.GenericFuncMap()

	localFuncMap := map[string]interface{}{
		"endpoint": func(m pgs.Method) string {
			var q rtkquerypb.MethodOptions
			ok, err := m.Extension(rtkquerypb.E_Query, &q)
			if err != nil {
				return err.Error()
			}
			if ok {
				return q.EndpointType.String()
			}
			return ""
		},
		"lowerFirst": func(s pgs.Name) string {
			for i, v := range s {
				return string(unicode.ToLower(v)) + string(s[i+1:])
			}
			return ""
		},
		"fname": func(s pgs.Name) string {
			return strings.TrimSuffix(filepath.Base(s.String()), filepath.Ext(s.String()))
		},
		"fnamelc": func(s pgs.Name) string {
			fn := strings.TrimSuffix(filepath.Base(s.String()), filepath.Ext(s.String()))
			return pgs.Name(strings.ReplaceAll(fn, "-", ".")).LowerCamelCase().String()
		},
	}

	tpl.Funcs(mergeFuncMaps(sprigFuncMap, localFuncMap))

	template.Must(tpl.Parse(fileTpl))

	storeFileData := &storeFile{}

	for _, f := range targets {
		m.Log("target", f.Name().String())
		if len(f.Services()) == 0 {
			continue
		}
		addStoreFileData(storeFileData, f)
		m.Push(f.Name().String())
		out := strings.TrimSuffix(f.Name().String(), "proto")
		out = out + "api.ts"
		imp, names := mkImports(f)
		m.AddGeneratorTemplateFile(out, tpl, map[string]interface{}{
			"file":    f,
			"imports": imp,
			"names":   names,
		})
		m.Pop()
	}

	pth := ""
	for f := range targets {
		pth = filepath.Dir(f)
		break
	}

	m.AddGeneratorTemplateFile(pth+"/store.ts", template.Must(template.New("store").Parse(storeTpl)), storeFileData)

	return m.Artifacts()
}

func addStoreFileData(sd *storeFile, f pgs.File) {
	fn := strings.TrimSuffix(f.Name().String(), "proto") + "api"
	imp := strings.Builder{}
	imp.WriteString("import { ")
	for i, s := range f.Services() {
		sn := s.Name().LowerCamelCase()
		imp.WriteString(sn.String())
		if i != len(f.Services())-1 {
			imp.WriteByte(',')
		}
		imp.WriteByte(' ')
		sd.Reducers = append(sd.Reducers, fmt.Sprintf("[%s.reducerPath]: %s.reducer,", sn, sn))
		sd.Middlewares = append(sd.Middlewares, sn.String()+".middleware,")
	}
	imp.WriteString(fmt.Sprintf("} from './%s'", fn))
	sd.Imports = append(sd.Imports, imp.String())
}

func mkImports(f pgs.File) ([]string, map[string]string) {

	out := []string{}
	imports := map[string]map[pgs.Message]struct{}{}
	names := map[string]string{}
	uniqNames := map[string]struct{}{}

	objects := []whatAs{}
	for _, s := range f.Services() {
		objects = append(objects, whatAs{
			what: string(s.Name()) + "Client",
		})
		for _, m := range s.Methods() {
			if _, ok := imports[string(m.Input().File().Name())]; !ok {
				imports[string(m.Input().File().Name())] = map[pgs.Message]struct{}{}
			}
			imports[string(m.Input().File().Name())][m.Input()] = struct{}{}
			if _, ok := imports[string(m.Output().File().Name())]; !ok {
				imports[string(m.Output().File().Name())] = map[pgs.Message]struct{}{}
			}
			imports[string(m.Output().File().Name())][m.Output()] = struct{}{}
		}
	}
	fn := strings.TrimSuffix(string(f.Name()), ".proto")
	out = append(out, genImportStatement(objects, "./"+fn+".client"))

	for fn, msgs := range imports {
		fn = strings.TrimSuffix(fn, ".proto")
		objects := []whatAs{}
		for msg := range msgs {
			if _, ok := names[msg.FullyQualifiedName()]; !ok {
				name := msg.Name().String()
				for {
					if _, ok := uniqNames[name]; !ok {
						uniqNames[name] = struct{}{}
						break
					}
					name = name + "$"
				}
				names[msg.FullyQualifiedName()] = name
			}
			objects = append(objects, whatAs{
				what: msg.Name().String(),
				as:   names[msg.FullyQualifiedName()],
			})
		}
		out = append(out, genImportStatement(objects, "./"+fn))
	}

	return out, names
}

type whatAs struct {
	what string
	as   string
}

func genImportStatement(objects []whatAs, from string) string {
	imp := strings.Builder{}
	imp.WriteString("import { ")
	for i, s := range objects {
		imp.WriteString(s.what)
		if s.as != "" && s.what != s.as {
			imp.WriteString(" as " + s.as)
		}
		if i != len(objects)-1 {
			imp.WriteByte(',')
		}
		imp.WriteByte(' ')
	}
	imp.WriteString(fmt.Sprintf("} from '%s'", from))
	return imp.String()
}

func mergeFuncMaps(maps ...map[string]interface{}) map[string]interface{} {
	fm := make(map[string]interface{})
	for _, m := range maps {
		for k, v := range m {
			fm[k] = v
		}
	}
	return fm
}

var _ pgs.Module = (*Module)(nil)

const fileTpl = `// Code generated by protoc-gen-rtk-query. DO NOT EDIT.
// source: {{ .file.InputPath }}
{{- $fn := fname .file.Name}}
{{- $fnl := fnamelc .file.Name}}

import { createApi } from '@reduxjs/toolkit/query/react'
import { GrpcWebFetchTransport } from '@protobuf-ts/grpcweb-transport';
import { grpcBaseQuery } from '@greatliontech/protobuf-ts-rtk-query';
{{- range .imports}}
{{.}}
{{- end}}

const transport = new GrpcWebFetchTransport({
  baseUrl: "http://localhost:5080"
});

{{- range .file.Services }}
{{- $sn := .Name }}
{{- $snl := .Name.LowerCamelCase }}

const {{$snl}}Client = new {{$sn}}Client(transport)

// Define a service using a base URL and expected endpoints
export const {{$snl}} = createApi({
  reducerPath: '{{$snl}}',
  baseQuery: grpcBaseQuery(),
  endpoints: (builder) => ({
{{- range .Methods }}
{{- $mn := .Name }}
{{- $mnl := .Name | lowerFirst }}
    {{$mnl}}: builder.{{if eq (endpoint .) "QUERY"}}query{{else}}mutation{{end}}<{{index $.names .Output.FullyQualifiedName}}, {{index $.names .Input.FullyQualifiedName}}>({
      query: (req) => {{$snl}}Client.{{$mnl}}(req)
    }),
{{- end }}
  }),
})

// Export hooks for usage in function components, which are
// auto-generated based on the defined endpoints
export const {
{{- range .Methods }}
  use{{.Name}}{{if eq (endpoint .) ("QUERY")}}Query{{else}}Mutation{{end}},
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
