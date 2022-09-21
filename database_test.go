package sql_db_plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateMysqlConnection(t *testing.T) {
	t.Run("initialize success", func(t *testing.T) {
		builder := &DatabasePropertiesBuilder{}
		props := builder.WithHost("localhost").WithUsername("root").WithPassword("root").WithPort(3306).WithDriver("mysql").Build()

		database := NewDatabase(props)

		assert.NotNil(t, database)
		database.Close()
	})
}

func TestCreateSqliteConnection(t *testing.T) {
	t.Run("initialize success", func(t *testing.T) {
		builder := &DatabasePropertiesBuilder{}
		props := builder.WithDriver("sqlite").WithDB("test").Build()

		database := NewDatabase(props)

		assert.NotNil(t, database)

		database.Close()
	})
}
