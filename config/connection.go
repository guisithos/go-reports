package config

import (
	"context"
	"database/sql"

	"github.com/godror/godror"
)

var (
	db  *sql.DB
	err error
)

func init() {
	params := godror.ConnectionParams{
		Username:      "user",
		Password:      godror.NewPassword("password"),
		ConnectString: "localhost.66",
	}

	db, err = sql.Open("godror", params.StringWithPassword())
	if err != nil {
		panic(err)
	}

	if err = db.PingContext(context.Background()); err != nil {
		panic(err)
	}
}

func GetDB() (*sql.DB, error) {
	return db, err
}
