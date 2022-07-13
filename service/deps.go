package service

import (
	"../store"
	"database/sql"
)

type Dependencies struct {
	TinyUrlStore store.TinyUrl
}

func Init(db *sql.DB) (deps Dependencies) {
	deps.TinyUrlStore.DB = db
	return
}
