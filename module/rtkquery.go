package module

import (
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/Masterminds/sprig/v3"
	pgs "github.com/lyft/protoc-gen-star"
)

type Module struct {
	*pgs.ModuleBase
	params moduleParams
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

	params := parseParams(m.Parameters())
	m.params = params

	tpl := template.New("")

	sprigFuncMap := sprig.GenericFuncMap()

	localFuncMap := funcs(m)

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
		imp, names := getImportsAndMsgNames(f, params)
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
