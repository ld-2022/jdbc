package main

import (
	"github.com/jmoiron/sqlx"
)

var (
	dataSource *sqlx.DB
)

func SetDataSource(db *sqlx.DB) {
	dataSource = db
}
