package db

type DBConfig struct {
	Database string
}

func NewConn(db string) DBConfig {
	return DBConfig{
		Database: db,
	}
}

