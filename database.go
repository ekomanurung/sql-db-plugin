package sql_db_plugin

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

var MysqlDbConnectionString = "%s:%s@tcp(%s:%d)/%s?parseTime=true"
var PostgresSqlConnectionString = "host=%s user=%s dbname=%s password=%s port=%d sslmode=disable"
var SqliteConnectionString = "%s.db"

func GetDataSourceName(pattern string, props ...any) string {
	return fmt.Sprintf(pattern, props...)
}

func NewDatabase(props *DatabaseProperties, dialect gorm.Dialector) *sql.DB {
	var db *sql.DB

	conn, err := gorm.Open(dialect, &gorm.Config{})

	if err != nil {
		panic(fmt.Sprintf("Panic when initialize %s driver connection caused by: %+v\n", props.DriverName, err))
	}

	db, err = conn.DB()
	if err != nil {
		panic(fmt.Sprintf("Panic when get db driver %s caused by: %+v\n", props.DriverName, err))
	}

	db.SetMaxOpenConns(props.MaxOpenConnection)
	db.SetMaxIdleConns(props.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(props.MaxLifeTime) * time.Minute)

	return db
}
