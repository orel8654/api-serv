package database

import (
	"api/config"
	"context"
	"github.com/go-pg/pg/v10"
)

type NewConn struct {
	Conn *pg.DB
}

func NewDB(connect *pg.DB) *NewConn {
	return &NewConn{
		Conn: connect,
	}
}

func (db *NewConn) WriteRows(rows []config.DatabaseFields) error {
	_, err := db.Conn.Model(&rows).Insert()
	if err != nil {
		return err
	}
	return nil
}

func (db *NewConn) SelectRows() ([]config.DatabaseFields, error) {
	var result []config.DatabaseFields
	if err := db.Conn.Model(&result).Order("updated_at ASC").Select(); err != nil {
		return nil, err
	}
	return result, nil
}

func (db *NewConn) Ping() error {
	ctx := context.Background()
	if err := db.Conn.Ping(ctx); err != nil {
		return err
	}
	return nil
}