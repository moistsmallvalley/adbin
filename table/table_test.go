package table

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParse(t *testing.T) {
	tests := []struct {
		ddl    string
		tables []Table
	}{
		{
			ddl: "create table users(name text)",
			tables: []Table{{
				Name:    "users",
				Columns: []Column{{Name: "name", Type: TypeText}},
			}},
		},
		{
			ddl: "CREATE TABLE USERS(NAME TEXT)",
			tables: []Table{{
				Name:    "USERS",
				Columns: []Column{{Name: "NAME", Type: TypeText}},
			}},
		},
		{
			ddl: "create table users(id int, name text)",
			tables: []Table{{
				Name: "users",
				Columns: []Column{
					{Name: "id", Type: TypeInt},
					{Name: "name", Type: TypeText},
				},
			}},
		},
		{
			ddl: "create table users(id int not null)",
			tables: []Table{{
				Name:    "users",
				Columns: []Column{{Name: "id", Type: TypeInt, Required: true}},
			}},
		},
		{
			ddl: "create table users(id int); create table messages(text text)",
			tables: []Table{{
				Name:    "users",
				Columns: []Column{{Name: "id", Type: TypeInt}},
			}, {
				Name:    "messages",
				Columns: []Column{{Name: TypeText, Type: TypeText}},
			}},
		},
		{
			ddl: "create table users(id int primary key)",
			tables: []Table{{
				Name:    "users",
				Columns: []Column{{Name: "id", Type: TypeInt, Required: true, PrimaryKey: true}},
			}},
		},
		{
			ddl: "create table users(id int, primary key (id))",
			tables: []Table{{
				Name:    "users",
				Columns: []Column{{Name: "id", Type: TypeInt, Required: true, PrimaryKey: true}},
			}},
		},
		{
			ddl: "create table users(id int auto_increment)",
			tables: []Table{{
				Name:    "users",
				Columns: []Column{{Name: "id", Type: TypeInt, AutoIncrement: true}},
			}},
		},
	}
	for _, test := range tests {
		t.Run(test.ddl, func(t *testing.T) {
			tables, err := Parse(test.ddl)
			require.NoError(t, err)
			assert.Equal(t, test.tables, tables)
		})
	}
}
