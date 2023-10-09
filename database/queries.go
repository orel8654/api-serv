package database

import (
	"api/types"
	"fmt"
	"time"

	"github.com/fatih/structs"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Storage struct {
	s *sqlx.DB
}

func NewStorage(conf types.ConfDB) (*Storage, error) {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", conf.Username, conf.Password, conf.Database, conf.Host, conf.Port)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return nil, err
	}
	return &Storage{
		s: db,
	}, nil
}

func (s *Storage) WriteRow(data types.DataPost) error {
	dataWrite := types.DatabaseFields{
		CurrencyFrom: data.CurrencyFrom,
		CurrencyTo:   data.CurrencyTo,
		Well:         0,
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	query := `INSERT INTO newtable (currency_from, currency_to, well, updated_at) 
        	VALUES(:currency_from, :currency_to, :well, :updated_at)`
	_, err := s.s.NamedExec(query, dataWrite)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateWell(newData types.DataPut) error {
	query := `UPDATE newtable
			set well =$1
			where currency_to =$2`
	_, err := s.s.Exec(query, newData.Well, newData.CurrencyTo)
	if err != nil {
		return err
	}
	return nil
}

func (s *Storage) UpdateRows(newData types.CurrencyLatest) error {
	m := structs.Map(newData.Data)
	query := `UPDATE newtable
			set well =$1,
			updated_at =$2
			where currency_to =$3`

	for key := range m {
		_, err := s.s.Exec(query, m[key], time.Now().Format("2006-01-02 15:04:05"), key)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Storage) SelectRow(data types.DatabaseFields) (types.DatabaseFields, error) {
	query := `SELECT currency_from, currency_to, well, updated_at
			FROM newtable
			WHERE currency_from =$currency_from AND currency_to =$currency_to`
	res := types.DatabaseFields{}
	if err := s.s.Get(&res, query, data.CurrencyFrom, data.CurrencyTo); err != nil {
		return res, err
	}
	return res, nil
}

func (s *Storage) Exists(data types.DataPost) error {
	query := `SELECT currency_from, currency_to, well, updated_at
			FROM newtable
			WHERE currency_to =$1`
	var res types.DatabaseFields
	if err := s.s.Get(&res, query, data.CurrencyTo); err != nil {
		return err
	}
	return nil
}

func (s *Storage) SelectAll() ([]types.DatabaseFields, error) {
	query := `SELECT currency_from, currency_to, well, updated_at FROM newtable`
	var res []types.DatabaseFields
	if err := s.s.Select(&res, query); err != nil {
		return res, err
	}
	return res, nil
}
