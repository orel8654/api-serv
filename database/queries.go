package database

import (
	"api/config"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	s *sqlx.DB
}

func NewStorage(conf config.ConfDB) (*Storage, error) {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", conf.Username, conf.Password, conf.Database, conf.Host, conf.Port)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return nil, err
	}
	return &Storage{
		s: db,
	}, nil
}

func (s *Storage) WriteRow(data config.DatabaseFields) error {
	query := `INSERT INTO newtable(currency_from, currency_to, well, updated_at) 
        	VALUES(:currency_from, :currency_to, :well, :updated_at)`
	_, err := s.s.NamedExec(query, data)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) SelectRow(data config.DatabaseFields) (config.DatabaseFields, error) {
	query := `SELECT currency_from, currency_to, well, updated_at
			FROM newtable
			WHERE currency_from =$currency_from AND currency_to =$currency_to`
	res := config.DatabaseFields{}
	if err := s.s.Get(&res, query, data.CurrencyFrom, data.CurrencyTo); err != nil {
		return res, err
	}
	return res, nil
}

func (s *Storage) Exists(data config.DatabaseFields) (bool, error) {
	query := `SELECT currency_from, currency_to, well, updated_at
			FROM newtable
			WHERE currency_from =$currency_from AND currency_to =$currency_to`
	res := config.DatabaseFields{}
	if err := s.s.Get(&res, query, data.CurrencyFrom, data.CurrencyTo); err != nil {
		return false, err
	}
	return true, nil
}

func (s *Storage) SelectAll() (config.DatabaseFields, error) {
	query := `SELECT (currency_from, currency_to, well) FROM newtable`
	res := config.DatabaseFields{}
	if err := s.s.Select(&res, query); err != nil {
		return res, err
	}
	return res, nil
}
