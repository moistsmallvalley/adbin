package table

import (
	"strings"

	"github.com/pkg/errors"
)

type Type string

const (
	TypeUnknown = ""

	TypeInteger   = "integer"
	TypeInt       = "int"
	TypeSmallint  = "smallint"
	TypeTinyInt   = "tinyint"
	TypeMediumInt = "mediumint"
	TypeBitIng    = "bigint"
	TypeDecimal   = "decimal"
	TypeNumeric   = "numeric"
	TypeFloat     = "float"
	TypeDouble    = "double"
	TypeBit       = "bit"

	TypeDate      = "date"
	TypeDateTime  = "datetime"
	TypeTimestamp = "timestamp"
	TypeTime      = "time"
	TypeYear      = "year"

	TypeChar      = "char"
	TypeVarChar   = "varchar"
	TypeBinary    = "binary"
	TypeVarBinary = "varbinary"
	TypeBlob      = "blob"
	TypeText      = "text"
	TypeEnum      = "enum"
	TypeSet       = "set"

	TypeJson = "json"
)

var typeNameToType = map[string]Type{
	"integer":   TypeInteger,
	"int":       TypeInt,
	"smallint":  TypeSmallint,
	"tinyint":   TypeTinyInt,
	"mediumint": TypeMediumInt,
	"bigint":    TypeBitIng,
	"decimal":   TypeDecimal,
	"numeric":   TypeNumeric,
	"float":     TypeFloat,
	"double":    TypeDouble,
	"bit":       TypeBit,
	"date":      TypeDate,
	"datetime":  TypeDateTime,
	"timestamp": TypeTimestamp,
	"time":      TypeTime,
	"year":      TypeYear,
	"char":      TypeChar,
	"varchar":   TypeVarChar,
	"binary":    TypeBinary,
	"varbinary": TypeVarBinary,
	"blob":      TypeBlob,
	"text":      TypeText,
	"enum":      TypeEnum,
	"set":       TypeSet,
	"json":      TypeJson,
}

func makeType(parserTypeName string) (Type, error) {
	lower := strings.ToLower(parserTypeName)
	if typ, ok := typeNameToType[lower]; ok {
		return typ, nil
	}
	return TypeUnknown, errors.Errorf("unsupported type '%s'", parserTypeName)
}
