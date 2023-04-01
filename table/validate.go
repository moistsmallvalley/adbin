package table

import (
	"encoding/json"
	"time"

	"github.com/pkg/errors"
)

func ValidateRow(table Table, row Row) error {
	for name, value := range row {
		c, ok := table.ColumnByName(name)
		if !ok {
			return errors.Errorf("unknown column: %s", name)
		}
		if value == nil {
			if c.Required {
				return errors.Errorf("%s is reqiured and cannot be null", name)
			}
			continue
		}
		switch c.Type {
		case TypeInteger, TypeInt, TypeSmallInt, TypeTinyInt, TypeMediumInt, TypeBigInt, TypeDecimal, TypeNumeric, TypeFloat, TypeDouble, TypeBit, TypeYear:
			switch value.(type) {
			case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64:
			default:
				return errors.Errorf("%s is not number", name)
			}
		case TypeChar, TypeVarChar, TypeText, TypeEnum:
			switch value.(type) {
			case string:
			default:
				return errors.Errorf("%s is not string", name)
			}
		case TypeDate, TypeDateTime, TypeTimestamp, TypeTime:
			switch value := value.(type) {
			case string:
				if _, err := time.Parse(time.RFC3339, value); err != nil {
					return errors.Errorf("%s is not RFC3339 string", name)
				}
			default:
				return errors.Errorf("%s is not RFC3339 string", name)
			}
		case TypeBinary, TypeVarBinary, TypeBlob:
			switch value.(type) {
			case []byte:
			default:
				return errors.Errorf("%s is not bytes", name)
			}
		case TypeSet:
			return errors.New("set columns not supported yet")
		case TypeJson:
			switch value := value.(type) {
			case string:
				var x any
				if err := json.Unmarshal([]byte(value), &x); err != nil {
					return errors.Errorf("%s is not JSON string", name)
				}
			default:
				return errors.Errorf("%s is not JSON string", name)
			}
		default:
			return errors.Errorf("unsupported column type: %s", c.Type)
		}
	}
	return nil
}
