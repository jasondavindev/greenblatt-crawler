package requester

import (
	"strings"

	"github.com/jasondavindev/statusinvest-crawler/config"
)

// FieldProperties Estrutura do parâmetro da request http
type FieldProperties map[string]float32

// RequestParam Estrutura final do data filtro da request http
type RequestParam map[string]interface{}

// ParseFiltersToParams Converte os filtros para formato de parametros
func ParseFiltersToParams(f *config.Filters) RequestParam {
	reqParam := make(RequestParam)

	for k, v := range *f {
		if isExcludedField(k) {
			reqParam[k] = v
			continue
		}

		parsedValues := parsenInterfaceToFloat32Slice(v)

		param := parseFilterToParam(k, parsedValues)

		if param != nil {
			reqParam[k] = *param
		} else {
			reqParam[k] = v
		}
	}

	return reqParam
}

func isExcludedField(key string) bool {
	for _, f := range GetExludedFields() {
		if f == strings.ToLower(key) {
			return true
		}
	}
	return false
}

func parseFilterToParam(key string, values []float32) *FieldProperties {
	prop := make(FieldProperties)

	prop["Item1"] = values[0]

	if len(values) == 2 {
		prop["Item2"] = values[1]
	}

	return &prop
}

// GetExludedFields Campos que não devem ser convertidos
func GetExludedFields() []string {
	return []string{"my_range", "segment", "subsector", "sector"}
}

func parsenInterfaceToFloat32Slice(i interface{}) []float32 {
	switch t := i.(type) {
	case []interface{}:
		sl := []float32{}

		for _, v := range t {
			switch vt := v.(type) {
			case float64:
				sl = append(sl, float32(vt))
			case float32:
				sl = append(sl, vt)
			case int:
				sl = append(sl, float32(vt))
			default:
				sl = append(sl, 0)
			}
		}

		return sl
	}

	return []float32{}
}
