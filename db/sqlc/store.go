package db

import (
	"context"
	"database/sql"
)

type Store interface {
	TrasferTx(ctx context.Context, arg TrasferTxParams) (TrasferTxResult, error)
	Querier
}

type SQLStore struct {
	db *sql.DB
	*Queries
}

func NewStore(db *sql.DB) Store {
	return &SQLStore{
		db:      db,
		Queries: New(db),
	}
}
