package greenblatt

import "log"

// GetOrZero Retorna o valor ou zero
func GetOrZero(v interface{}, e interface{}) float64 {
	if e != nil {
		log.Fatal(e)
	}

	return GetFloat64Value(v)
}

// GetFloat64Value Retorna valor do tipo float64
func GetFloat64Value(i interface{}) float64 {
	switch v := i.(type) {
	case float64:
		return v
	case float32:
		return float64(v)
	case int:
		return float64(v)
	default:
		return 0
	}
}
