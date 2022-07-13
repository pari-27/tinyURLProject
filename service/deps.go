package service

import (
	"database/sql"
	"github.com/pari-27/tinyURLProject/store"
)

type Dependencies struct {
	TinyUrlStore store.TinyUrl
}

func Init(db *sql.DB) (deps Dependencies) {
	deps.TinyUrlStore.DB = db
	return
}
