package module

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"
	"unicode"

	"github.com/Masterminds/sprig/v3"
	pgs "github.com/lyft/protoc-gen-star"

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
			ok, err := m.Extension(rtkquerypb.E_Endpoint, &q)
			if err != nil {
				return err.Error()
			}
			if ok {
				return q.Type.String()
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
		"hasProvidesTags": func(mth pgs.Method) (bool, error) {
			var q rtkquerypb.MethodOptions
			ok, err := mth.Extension(rtkquerypb.E_Endpoint, &q)
			if err != nil {
				return false, err
			}
			if ok {
				return q.ProvidesTags != nil, nil
			}
			return false, nil
		},
		"providesTags": func(mth pgs.Method) (string, error) {
			var q rtkquerypb.MethodOptions
			ok, err := mth.Extension(rtkquerypb.E_Endpoint, &q)
			if err != nil {
				return "", err
			}
			if ok {
				switch pt := q.ProvidesTags.(type) {
				case *rtkquerypb.MethodOptions_ProvidesList:
					itemsName := "items"
					if pt.ProvidesList.Items != nil {
						itemsName = *pt.ProvidesList.Items
					}
					pth, err := toJsPath(mth.Output(), itemsName)
					if err != nil {
						return "", err
					}
					return fmt.Sprintf("(result) => providesList(result?.%s, '%s')", pth, pt.ProvidesList.Tag), nil
				case *rtkquerypb.MethodOptions_ProvidesSpecific:
					idName := "id"
					if pt.ProvidesSpecific.Id != nil {
						idName = *pt.ProvidesSpecific.Id
					}
					return fmt.Sprintf("(result, error, arg) => [{ type: '%s', id: arg.%s }]", pt.ProvidesSpecific.Tag, idName), nil
				case *rtkquerypb.MethodOptions_ProvidesGeneric:
					return fmt.Sprintf("['%s']", pt.ProvidesGeneric), nil
				}
			}
			return "", nil
		},
		"hasInvalidatesTags": func(mth pgs.Method) bool {
			var q rtkquerypb.MethodOptions
			ok, err := mth.Extension(rtkquerypb.E_Endpoint, &q)
			if err != nil {
				m.Log("hasInvalidatesTags error", err)
				return false
			}
			if ok {
				return q.InvalidatesTags != nil
			}
			return false
		},
		"invalidatesTags": func(mth pgs.Method) string {
			var q rtkquerypb.MethodOptions
			ok, err := mth.Extension(rtkquerypb.E_Endpoint, &q)
			if err != nil {
				m.Log("invalidateTags error", err)
				return ""
			}
			if ok {
				switch pt := q.InvalidatesTags.(type) {
				case *rtkquerypb.MethodOptions_InvalidatesList:
					return fmt.Sprintf("(result, error, arg) => [{ type: '%s', id: 'LIST' }]", pt.InvalidatesList)
				case *rtkquerypb.MethodOptions_InvalidatesSpecific:
					idName := "id"
					if pt.InvalidatesSpecific.Id != nil {
						idName = *pt.InvalidatesSpecific.Id
						idName = strings.ReplaceAll(idName, ".", "?.")
					}
					return fmt.Sprintf("(result, error, arg) => [{ type: '%s', id: arg.%s }]", pt.InvalidatesSpecific.Tag, idName)
				case *rtkquerypb.MethodOptions_InvalidatesGeneric:
					return fmt.Sprintf("['%s']", pt.InvalidatesGeneric)
				}
			}
			return ""
		},
		"hasTags": func(svc pgs.Service) bool {
			var ep rtkquerypb.ServiceOptions
			ok, err := svc.Extension(rtkquerypb.E_Api, &ep)
			if err != nil {
				m.Log("hasTags error", err)
				return false
			}
			if ok {
				return ep.Tags != nil
			}
			return false
		},
		"tags": func(svc pgs.Service) string {
			var ep rtkquerypb.ServiceOptions
			ok, err := svc.Extension(rtkquerypb.E_Api, &ep)
			if err != nil {
				m.Log("hasTags error", err)
				return ""
			}
			if ok {
				return "['" + strings.Join(ep.Tags, "','") + "']"
			}
			return ""
		},
	}

	tpl.Funcs(mergeFuncMaps(sprigFuncMap, localFuncMap))

	template.Must(tpl.Parse(fileTpl))

	storeFileData := &storeFile{}

	for _, f := range targets {
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

	objects := []tsImport{}
	for _, s := range f.Services() {
		objects = append(objects, tsImport{
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

	ffn := strings.TrimSuffix(string(f.Name()), ".proto")
	fn := filepath.Base(ffn)
	dir := filepath.Dir(ffn)

	out = append(out, genImportStatement(objects, "./"+fn+".client"))

	for ifn, msgs := range imports {
		ifn = strings.TrimSuffix(ifn, ".proto")
		objects := []tsImport{}
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
			objects = append(objects, tsImport{
				what: msg.Name().String(),
				as:   names[msg.FullyQualifiedName()],
			})
		}
		out = append(out, genImportStatement(objects, genImportFileName(dir, ifn)))
	}

	return out, names
}

func genImportFileName(dir, fn string) string {
	fdir := filepath.Dir(fn)
	if dir == fdir {
		return "./" + filepath.Base(fn)
	}
	if dir == "." {
		return "./" + fn
	}
	sb := strings.Builder{}
	for range strings.Split(dir, "/") {
		sb.WriteString("../")
	}
	sb.WriteString(fn)
	return sb.String()
}

type tsImport struct {
	what string
	as   string
}

func genImportStatement(imports []tsImport, from string) string {
	imp := strings.Builder{}
	imp.WriteString("import { ")
	for i, s := range imports {
		imp.WriteString(s.what)
		if s.as != "" && s.what != s.as {
			imp.WriteString(" as " + s.as)
		}
		if i != len(imports)-1 {
			imp.WriteByte(',')
		}
		imp.WriteByte(' ')
	}
	imp.WriteString(fmt.Sprintf("} from '%s'", from))
	return imp.String()
}

func toJsPath(msg pgs.Message, pth string) (string, error) {
	parts := strings.Split(pth, ".")
	jsPath := []string{}

outer:
	for i, part := range parts {
		for _, fld := range msg.Fields() {
			if fld.Name().String() == part {
				jsPath = append(jsPath, *fld.Descriptor().JsonName)
				if fld.Type().ProtoType() == pgs.MessageT {
					msg = fld.Type().Embed()
					continue outer
				}
				// TODO: better error handling
				if i != len(parts)-1 {
					return "", fmt.Errorf("invalid path %q for message %q", pth, msg.FullyQualifiedName())
				}
				break outer
			}
			continue
		}
		return "", fmt.Errorf("message field %q not found in message %q", part, msg.FullyQualifiedName())
	}

	return strings.Join(jsPath, "?."), nil
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
import { GrpcWebFetchTransport, GrpcWebOptions } from '@protobuf-ts/grpcweb-transport';
import { grpcBaseQuery, providesList } from '@greatliontech/protobuf-ts-rtk-query';
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
{{- $mnl := .Name | lowerFirst }}
    {{$mnl}}: builder.{{if eq (endpoint .) "QUERY"}}query{{else}}mutation{{end}}<{{index $.names .Output.FullyQualifiedName}}, {{index $.names .Input.FullyQualifiedName}}>({
      query: (req) => {{$snl}}Client.{{$mnl}}(req),
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
