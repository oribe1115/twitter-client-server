package model

import (
	"github.com/antonlindstrom/pgstore"
)

var (
	store *pgstore.PGStore
)

// セッションを作成
func storeForSession() (*pgstore.PGStore, error) {
	_store, err := pgstore.NewPGStore(databaseURL, []byte("secret-key"))

	if err != nil {
		return nil, err
	}
	store = _store

	return store, nil
}
