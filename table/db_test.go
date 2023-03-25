package table

import (
	"database/sql"
	"testing"

	_ "github.com/go-sql-driver/mysql"
	"github.com/moistsmallvalley/adbin/testdb"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testDBUser = "root"
	testDBPass = "example"
	testDBName = "tabletest"
)

func initTestDB(ddls ...string) (*sql.DB, error) {
	return testdb.InitTestDB(testDBUser, testDBPass, testDBName, ddls...)
}

func TestListDBTable(t *testing.T) {
	db, err := initTestDB(
		`
CREATE TABLE users (
  id INT AUTO_INCREMENT PRIMARY KEY,
  username VARCHAR(255) NOT NULL UNIQUE
)`,
		`
CREATE TABLE mails (
  id INT AUTO_INCREMENT PRIMARY KEY,
  user_id INT NOT NULL,
  text TEXT NOT NULL,
  FOREIGN KEY (user_id) REFERENCES users (id)
)`,
	)
	require.NoError(t, err)
	defer db.Close()

	tables, err := ListDBTables(db)
	require.NoError(t, err)

	assert.Equal(t, []Table{
		{
			Name: "mails",
			Columns: []Column{
				{
					Name:          "id",
					Type:          TypeInt,
					Required:      true,
					PrimaryKey:    true,
					AutoIncrement: true,
				},
				{
					Name:     "user_id",
					Type:     TypeInt,
					Required: true,
				},
				{
					Name:     "text",
					Type:     TypeText,
					Required: true,
				},
			},
		},
		{
			Name: "users",
			Columns: []Column{
				{
					Name:          "id",
					Type:          TypeInt,
					Required:      true,
					PrimaryKey:    true,
					AutoIncrement: true,
				},
				{
					Name:     "username",
					Type:     TypeVarChar,
					Required: true,
				},
			},
		},
	}, tables)
}
