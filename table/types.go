package table

import (
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/pkg/errors"
)

type Type string

const (
	TypeUnknown = ""

	TypeInteger   = "integer"
	TypeInt       = "int"
	TypeSmallInt  = "smallint"
	TypeTinyInt   = "tinyint"
	TypeMediumInt = "mediumint"
	TypeBigInt    = "bigint"
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
	"smallint":  TypeSmallInt,
	"tinyint":   TypeTinyInt,
	"mediumint": TypeMediumInt,
	"bigint":    TypeBigInt,
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

func (t Type) Parse(s string) (any, error) {
	switch t {
	case TypeInteger, TypeInt:
		return strconv.ParseInt(s, 10, 32)
	case TypeTinyInt:
		return strconv.ParseInt(s, 10, 8)
	case TypeSmallInt:
		return strconv.ParseInt(s, 10, 16)
	case TypeMediumInt:
		return strconv.ParseInt(s, 10, 24)
	case TypeBigInt:
		return strconv.ParseInt(s, 10, 64)
	case TypeFloat:
		return strconv.ParseFloat(s, 32)
	case TypeDecimal, TypeNumeric, TypeDouble:
		return strconv.ParseFloat(s, 64)
	case TypeBit:
		return strconv.ParseUint(s, 10, 8)
	case TypeChar, TypeVarChar, TypeText:
		return s, nil
	case TypeEnum:
		return s, nil // TODO: check enum value
	case TypeDate:
		return time.Parse("2006-01-02", s)
	case TypeDateTime:
		return time.Parse(time.RFC3339, s)
	case TypeTimestamp:
		return strconv.ParseFloat(s, 64)
	case TypeTime:
		return time.Parse("15:04:05", s)
	case TypeBinary, TypeVarBinary, TypeBlob:
		return base64.StdEncoding.DecodeString(s)
	case TypeSet:
		return nil, errors.New("set columns not supported yet")
	case TypeJson:
		return s, nil
	default:
		return nil, errors.Errorf("unsupported column type: %s", t)
	}
}
