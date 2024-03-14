package parser

const (
	INT8          = "int8"
	INT16         = "int16"
	INT32         = "int32"
	INT64         = "int64"
	UINT8         = "uint8"
	UINT16        = "uint16"
	UINT32        = "uint32"
	UINT64        = "uint64"
	INT           = "int"
	UINT          = "uint"
	RUNE          = "rune"
	BYTE          = "byte"
    BYTE_SLICE    = "[]byte"
	FLOAT32       = "float32"
	FLOAT64       = "float64"
	COMPLEX64     = "complex64"
	COMPLEX128    = "complex128"
	BOOL          = "bool"
	STRING        = "string"
	ERROR         = "error"
	TIME_DURATION = "time.Duration"
)

var TypesValidationMap = map[string]bool{
	"int8":          true,
	"int16":         true,
	"int32":         true,
	"int64":         true,
	"uint8":         true,
	"uint16":        true,
	"uint32":        true,
	"uint64":        true,
	"int":           true,
	"uint":          true,
	"rune":          true,
	"byte":          true,
    "[]byte":        true,
	"float32":       true,
	"float64":       true,
	"complex64":     true,
	"complex128":    true,
	"bool":          true,
	"string":        true,
	"error":         true,
	"time.Duration": true,
}

func GetType(s string) string {
	if TypesValidationMap[s] {
		return s
	}
	return "string"
}
