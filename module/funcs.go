package module

import (
	"fmt"
	"path/filepath"
	"strings"
	"unicode"

	rtkquerypb "github.com/greatliontech/protoc-gen-rtk-query/proto/rtkquery"
	pgs "github.com/lyft/protoc-gen-star"
)

func funcs(m pgs.ModuleBase) map[string]interface{} {
	return map[string]interface{}{
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

}
