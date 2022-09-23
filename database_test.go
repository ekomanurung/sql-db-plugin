package sql_db_plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/driver/sqlite"
)

func TestCreateMysqlConnection(t *testing.T) {
	t.Run("initialize success", func(t *testing.T) {
		builder := &DatabasePropertiesBuilder{}
		props := builder.WithHost("localhost").WithUsername("root").WithPassword("root").WithPort(3306).WithDriver("mysql").Build()

		database := NewDatabase(props, mysql.New(mysql.Config{
			DSN: GetDataSourceName(MysqlDbConnectionString,
				props.Username, props.Password)}))

		assert.NotNil(t, database)
		database.Close()
	})
}

func TestCreateSqliteConnection(t *testing.T) {
	t.Run("initialize success", func(t *testing.T) {
		builder := &DatabasePropertiesBuilder{}
		props := builder.WithDriver("sqlite").WithDB("test").Build()

		database := NewDatabase(props, sqlite.Open(GetDataSourceName(SqliteConnectionString, props.DatabaseName)))

		assert.NotNil(t, database)

		database.Close()
	})
}
