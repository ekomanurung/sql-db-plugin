package sql_db_plugin

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAddDatabaseProperties(t *testing.T) {
	t.Run("create props success", func(t *testing.T) {
		properties := NewDatabaseProperties("eko", "test", "localhost", 3306, "test", "mysql")
		assert.NotNil(t, properties)
	})

	t.Run("builder - success", func(t *testing.T) {
		dbBuilder := &DatabasePropertiesBuilder{}
		props := dbBuilder.WithUsername("root").WithPassword("root").WithDriver("mysql").WithHost("localhost").Build()
		assert.NotNil(t, props)
		assert.Equal(t, "mysql", props.DriverName)
	})

	t.Run("builder - empty host", func(t *testing.T) {
		dbBuilder := &DatabasePropertiesBuilder{}
		assertPanic(t, func() {
			dbBuilder.WithUsername("root").WithPassword("root").WithDriver("mysql").WithHost("").Build()
		})
	})
}

func assertPanic(t *testing.T, f func()) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()
	f()
}
