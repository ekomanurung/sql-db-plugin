package sql_db_plugin

type DatabaseProperties struct {
	Username          string
	Password          string
	Host              string
	Port              int
	DatabaseName      string
	DriverName        string
	MaxOpenConnection int
	MaxIdleConnection int
	MaxLifeTime       int
}

func NewDatabaseProperties(username string, password string, host string, port int, databaseName string, driverName string) *DatabaseProperties {
	return &DatabaseProperties{Username: username, Password: password, Host: host, Port: port, DatabaseName: databaseName, DriverName: driverName}
}

func (db *DatabaseProperties) SetMaxOpenConnection(maxOpenConn int) {
	db.MaxOpenConnection = maxOpenConn
}

func (db *DatabaseProperties) SetMaxIdleConnection(maxIdleConn int) {
	db.MaxIdleConnection = maxIdleConn
}

func (db *DatabaseProperties) SetMaxLifeTime(maxLifeTime int) {
	db.MaxIdleConnection = maxLifeTime
}

type DatabasePropertiesBuilder struct {
	Username          string
	Password          string
	Host              string
	Port              int
	DatabaseName      string
	DriverName        string
	MaxOpenConnection int
	MaxIdleConnection int
	MaxLifeTime       int

	DatabaseProperties *DatabaseProperties
}

type DBBuilder interface {
	WithUsername(username string)
	WithPassword(pwd string)
	WithHost(host string)
	WithPort(port int)
	WithDB(db string)
	WithDriver(driver string)
	WithMaxOpenConnection(maxOpenConn int)
	WithMaxIdleConnection(maxIdleConn int)
	WithMaxLifeTime(maxLifeTime int)
	Build() DatabaseProperties
}

func (d *DatabasePropertiesBuilder) WithUsername(username string) *DatabasePropertiesBuilder {
	d.Username = username
	return d
}

func (d *DatabasePropertiesBuilder) WithPassword(pwd string) *DatabasePropertiesBuilder {
	d.Password = pwd
	return d
}

func (d *DatabasePropertiesBuilder) WithHost(host string) *DatabasePropertiesBuilder {
	if host == "" {
		panic("database host can't be empty")
	}

	d.Host = host
	return d
}

func (d *DatabasePropertiesBuilder) WithPort(port int) *DatabasePropertiesBuilder {
	if port < 1000 {
		port = 3000
	}

	d.Port = port
	return d
}

func (d *DatabasePropertiesBuilder) WithDB(db string) *DatabasePropertiesBuilder {
	if db == "" {
		panic("database name can't be empty")
	}

	d.DatabaseName = db
	return d
}

func (d *DatabasePropertiesBuilder) WithDriver(driver string) *DatabasePropertiesBuilder {
	if driver == "" {
		panic("database driver can't be empty")
	}

	d.DriverName = driver
	return d
}

func (d *DatabasePropertiesBuilder) WithMaxOpenConnection(maxOpenConn int) *DatabasePropertiesBuilder {
	if maxOpenConn < 1 {
		maxOpenConn = 10
	}

	d.MaxOpenConnection = maxOpenConn
	return d
}

func (d *DatabasePropertiesBuilder) WithMaxIdleConnection(maxIdleConn int) *DatabasePropertiesBuilder {
	if maxIdleConn < 1 {
		maxIdleConn = 10
	}

	d.MaxIdleConnection = maxIdleConn
	return d
}

func (d *DatabasePropertiesBuilder) WithMaxLifeTime(maxLifeTime int) *DatabasePropertiesBuilder {
	if maxLifeTime < 1 {
		maxLifeTime = 10
	}

	d.MaxLifeTime = maxLifeTime
	return d
}

func (d *DatabasePropertiesBuilder) Build() *DatabaseProperties {
	return &DatabaseProperties{
		Username:          d.Username,
		Password:          d.Password,
		Host:              d.Host,
		Port:              d.Port,
		DatabaseName:      d.DatabaseName,
		DriverName:        d.DriverName,
		MaxOpenConnection: d.MaxOpenConnection,
		MaxIdleConnection: d.MaxIdleConnection,
		MaxLifeTime:       d.MaxLifeTime,
	}
}
