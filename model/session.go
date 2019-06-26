package model

import (
	"github.com/antonlindstrom/pgstore"
)

var (
	Store *pgstore.PGStore
)

// セッションを作成
func storeForSession() (*pgstore.PGStore, error) {
	_store, err := pgstore.NewPGStore(databaseURL, []byte("secret-key"))

	if err != nil {
		return nil, err
	}
	Store = _store

	return Store, nil
}
