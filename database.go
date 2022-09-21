package sql_db_plugin

import (
	"database/sql"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var MysqlDbConnectionString = "%s:%s@tcp(%s:%d)/%s?parseTime=true"
var PostgresSqlConnectionString = "host=%s user=%s dbname=%s password=%s port=%d sslmode=disable"

func getDataSourceName(pattern string, props ...any) string {
	return fmt.Sprintf(pattern, props...)
}

func NewDatabase(props *DatabaseProperties) *sql.DB {
	var db *sql.DB
	switch props.DriverName {
	case "mysql":
		conn, err := gorm.Open(mysql.New(mysql.Config{
			DSN: getDataSourceName(MysqlDbConnectionString,
				props.Username, props.Password,
				props.Host, props.Port, props.DatabaseName),
		}), &gorm.Config{})

		if err != nil {
			panic(fmt.Sprintf("Panic when initialize %s driver connection caused by: %+v\n", props.DriverName, err))
		}

		db, err = conn.DB()
		if err != nil {
			panic(fmt.Sprintf("Panic when get db driver %s caused by: %+v\n", props.DriverName, err))
		}
	case "sqlite":
		conn, err := gorm.Open(sqlite.Open(fmt.Sprintf("%s.db", props.DatabaseName)), &gorm.Config{})
		if err != nil {
			panic(fmt.Sprintf("Panic when initialize Gorm with driver %s caused by: %+v\n", props.DriverName, err))
		}

		db, err = conn.DB()
		if err != nil {
			panic(fmt.Sprintf("Panic when get db driver %s caused by: %+v\n", props.DriverName, err))
		}
	case "postgresql":
		conn, err := gorm.Open(postgres.New(postgres.Config{
			DSN: getDataSourceName(PostgresSqlConnectionString,
				props.Host, props.Username,
				props.DatabaseName, props.Password, props.Port),
		}), &gorm.Config{})

		if err != nil {
			panic(fmt.Sprintf("Panic when initialize %s driver connection caused by: %+v\n", props.DriverName, err))
		}

		db, err = conn.DB()
		if err != nil {
			panic(fmt.Sprintf("Panic when get db driver %s caused by: %+v\n", props.DriverName, err))
		}
	default:
		panic("Unimplemented database driver..")
	}

	db.SetMaxOpenConns(props.MaxOpenConnection)
	db.SetMaxIdleConns(props.MaxIdleConnection)
	db.SetConnMaxLifetime(time.Duration(props.MaxLifeTime) * time.Minute)

	return db
}
