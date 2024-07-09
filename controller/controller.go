package controller

import "database/sql"

type Controllers struct {
	db *sql.DB
}

func NewController(db *sql.DB) *Controllers {
	return &Controllers{
		db: db,
	}
}
