package dbmysql

import (
	"backend/src/config"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectMySql() (*sql.DB, error) {
	db, err := sql.Open("mysql", config.ServerAddress)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		db.Close()
		return nil, err
	}
	return db, nil
}
