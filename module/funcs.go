package module

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	rtkquerypb "github.com/greatliontech/protoc-gen-rtk-query/proto/rtkquery"
	pgs "github.com/lyft/protoc-gen-star"
)

func funcs(mod *Module) map[string]interface{} {
	return map[string]interface{}{
		"endpoint": func(m pgs.Method) (string, error) {
			var q rtkquerypb.MethodOptions
			ok, err := m.Extension(rtkquerypb.E_Endpoint, &q)
			if err != nil {
				return "", err
			}
			if ok {
				if q.GetType() == rtkquerypb.EndpointType_QUERY {
					return "query", nil
				}
				return "mutation", nil
			}
			if mod.params.WithAIPStandardMethods {
				methName := m.Name().String()
				if strings.HasPrefix(methName, "Get") || strings.HasPrefix(methName, "List") {
					return "query", nil
				}
				if strings.HasPrefix(methName, "Create") || strings.HasPrefix(methName, "Update") || strings.HasPrefix(methName, "Delete") {
					return "mutation", nil
				}
			}
			return "", fmt.Errorf("cannot determine endpoint type for %s.%s", m.Name().String(), m.Service().Name().String())
		},
		"lowerFirst": func(s string) string {
			for i, v := range s {
				return string(unicode.ToLower(v)) + s[i+1:]
			}
			return ""
		},
		"upperFirst": func(s string) string {
			for i, v := range s {
				return string(unicode.ToUpper(v)) + s[i+1:]
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
					return fmt.Sprintf("(_1, _2, arg) => [{ type: '%s', id: arg.%s }]", pt.ProvidesSpecific.Tag, idName), nil
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
				mod.Log("hasInvalidatesTags error", err)
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
				mod.Log("invalidateTags error", err)
				return ""
			}
			if ok {
				switch pt := q.InvalidatesTags.(type) {
				case *rtkquerypb.MethodOptions_InvalidatesList:
					return fmt.Sprintf("() => [{ type: '%s', id: 'LIST' }]", pt.InvalidatesList)
				case *rtkquerypb.MethodOptions_InvalidatesSpecific:
					idName := "id"
					if pt.InvalidatesSpecific.Id != nil {
						idName = *pt.InvalidatesSpecific.Id
						idName = strings.ReplaceAll(idName, ".", "?.")
					}
					return fmt.Sprintf("(_1, _2, arg) => [{ type: '%s', id: arg.%s }]", pt.InvalidatesSpecific.Tag, idName)
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
				mod.Log("hasTags error", err)
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
				mod.Log("hasTags error", err)
				return ""
			}
			if ok {
				return "['" + strings.Join(ep.Tags, "','") + "']"
			}
			return ""
		},
		"withMetadata": func() bool {
			return mod.params.WithMetadata
		},
		"importName": func(m pgs.Message, names map[string]string) (string, error) {
			if n, ok := names[m.FullyQualifiedName()]; ok {
				return n, nil
			}
			return "", fmt.Errorf("cannot find input message name for %s in %s", m.Name().String(), m.File().Name())
		},
	}
}
