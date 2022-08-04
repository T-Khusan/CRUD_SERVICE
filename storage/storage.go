package storage

import (
	"crud_service/storage/postgres"

	"github.com/jmoiron/sqlx"
)

type StorageI interface {
	Crud() postgres.CrudI
}

type Store struct {
	db   *sqlx.DB
	crud postgres.CrudI
}

func NewStoragePg(db *sqlx.DB) StorageI {
	return &Store{
		db:   db,
		crud: postgres.NewCrudRepo(db),
	}
}

func (s *Store) Crud() postgres.CrudI {
	if s.crud == nil {
		s.crud = postgres.NewCrudRepo(s.db)
	}
	return s.crud
}
