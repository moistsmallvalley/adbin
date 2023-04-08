package table

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestValidateRow(t *testing.T) {
	t.Run("passes valid int", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "id", Type: TypeInt}},
		}
		row := Row{"id": 123}
		assert.NoError(t, ValidateRow(table, row))
	})

	t.Run("fails invalid int", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "id", Type: TypeInt}},
		}
		row := Row{"id": "123"}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("passes valid string", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "name", Type: TypeVarChar}},
		}
		row := Row{"name": "foo"}
		assert.NoError(t, ValidateRow(table, row))
	})

	t.Run("fails invalid string", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "name", Type: TypeVarChar}},
		}
		row := Row{"name": 123}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("passes valid datetime", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "created_at", Type: TypeDateTime}},
		}
		row := Row{"created_at": time.Date(2000, 1, 2, 10, 20, 30, 0, time.UTC)}
		assert.NoError(t, ValidateRow(table, row))
	})

	t.Run("fails invalid datetime", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "created_at", Type: TypeDateTime}},
		}
		row := Row{"created_at": 123}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("fails invalid datetime format", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "created_at", Type: TypeDateTime}},
		}
		row := Row{"created_at": "2000-01-02 10:20:30Z"}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("passes valid binary", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "data", Type: TypeBinary}},
		}
		row := Row{"data": []byte("test")}
		assert.NoError(t, ValidateRow(table, row))
	})

	t.Run("fails invalid binary", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "data", Type: TypeBinary}},
		}
		row := Row{"data": "test"}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("fails unnsuppored set", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "data", Type: TypeSet}},
		}
		row := Row{"data": "test"}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("passes valid json", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "data", Type: TypeJson}},
		}
		row := Row{"data": `{"foo": "bar"}`}
		assert.NoError(t, ValidateRow(table, row))
	})

	t.Run("fails invalid json", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "data", Type: TypeJson}},
		}
		row := Row{"data": 123}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("fails invalid json format", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "data", Type: TypeJson}},
		}
		row := Row{"data": "test"}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("passes nil unrequired column", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "id", Type: TypeInt, Required: false}},
		}
		row := Row{"id": nil}
		assert.NoError(t, ValidateRow(table, row))
	})

	t.Run("fails nil for required column", func(t *testing.T) {
		table := Table{
			Columns: []Column{{Name: "id", Type: TypeInt, Required: true}},
		}
		row := Row{"id": nil}
		assert.Error(t, ValidateRow(table, row))
	})

	t.Run("fails unknown column", func(t *testing.T) {
		table := Table{}
		row := Row{"id": 123}
		assert.Error(t, ValidateRow(table, row))
	})
}
