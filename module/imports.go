package module

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/hashicorp/go-set"
	pgs "github.com/lyft/protoc-gen-star"
)

type tsImport struct {
	What string
	As   string
}

func getImportsAndMsgNames(f pgs.File, params moduleParams) ([]string, map[string]string) {

	imports := []string{}

	objects := []tsImport{}
	for _, svc := range f.Services() {
		objects = append(objects, tsImport{
			What: string(svc.Name()) + "Client",
		})
	}

	imports = append(imports, genImportStatement(objects, genClientImportFileName(f, params)))

	// map msg fully qualified name to unique name
	names := map[string]string{}

	// map which imported source file contains which messages
	importFiles := getImportFiles(f.Services())

	// set to track unique names for imports. uses $ suffix to avoid name collisions
	uniqNames := set.New[string](1)
	getUniqName := func(name string) string {
		for {
			if !uniqNames.Contains(name) {
				uniqNames.Insert(name)
				break
			}
			name = name + "$"
		}
		return name
	}

	for file, msgs := range importFiles {

		objects := []tsImport{}

		for _, msg := range msgs.Slice() {

			if _, ok := names[msg.FullyQualifiedName()]; !ok {
				names[msg.FullyQualifiedName()] = getUniqName(msg.Name().String())
			}

			objects = append(objects, tsImport{
				What: msg.Name().String(),
				As:   names[msg.FullyQualifiedName()],
			})
		}
		imports = append(imports, genImportStatement(objects, genImportFileName(f, file, params)))
	}

	return imports, names
}

// returns a map of file to set of messages in that file from method imputs and outputs
func getImportFiles(svcs []pgs.Service) map[pgs.File]*set.Set[pgs.Message] {
	importFiles := map[pgs.File]*set.Set[pgs.Message]{}
	for _, svc := range svcs {
		for _, m := range svc.Methods() {
			s, ok := importFiles[m.Input().File()]
			if !ok {
				s = set.New[pgs.Message](1)
				importFiles[m.Input().File()] = s
			}
			s.Insert(m.Input())
			s, ok = importFiles[m.Output().File()]
			if !ok {
				s = set.New[pgs.Message](1)
				importFiles[m.Output().File()] = s
			}
			s.Insert(m.Output())
		}
	}
	return importFiles
}

func genClientImportFileName(currentFile pgs.File, params moduleParams) string {

	fn := strings.TrimSuffix(currentFile.Name().String(), ".proto")

	for k, v := range params.Imports {
		if strings.HasPrefix(currentFile.Package().ProtoName().String(), k) {
			p := filepath.Join(v, fn)
			if params.AddPbSuffix {
				return p + "_pb.client"
			}
			return p + ".client"
		}
	}
	if params.AddPbSuffix {
		return "./" + fn + "_pb.client"
	}
	return "./" + fn + ".client"
}

func genImportFileName(currentFile, importedFile pgs.File, params moduleParams) string {

	dir := filepath.Dir(strings.TrimSuffix(currentFile.Name().String(), ".proto"))
	fn := strings.TrimSuffix(importedFile.Name().String(), ".proto")

	for k, v := range params.Imports {
		if strings.HasPrefix(importedFile.Package().ProtoName().String(), k) {
			p := filepath.Join(v, fn)
			if params.AddPbSuffix {
				return p + "_pb"
			}
			return p
		}
	}

	fdir := filepath.Dir(fn)
	if dir == fdir {
		if params.AddPbSuffix {
			return "./" + filepath.Base(fn) + "_pb"
		}
		return "./" + filepath.Base(fn)
	}
	if dir == "." {
		if params.AddPbSuffix {
			return "./" + fn + "_pb"
		}
		return "./" + fn
	}
	sb := strings.Builder{}
	for range strings.Split(dir, "/") {
		sb.WriteString("../")
	}
	sb.WriteString(fn)
	if params.AddPbSuffix {
		return sb.String() + "_pb"
	}
	return sb.String()
}

func genImportStatement(imports []tsImport, from string) string {
	imp := strings.Builder{}
	imp.WriteString("import { ")
	for i, s := range imports {
		imp.WriteString(s.What)
		if s.As != "" && s.What != s.As {
			imp.WriteString(" as " + s.As)
		}
		if i != len(imports)-1 {
			imp.WriteByte(',')
		}
		imp.WriteByte(' ')
	}
	imp.WriteString(fmt.Sprintf("} from '%s'", from))
	return imp.String()
}
