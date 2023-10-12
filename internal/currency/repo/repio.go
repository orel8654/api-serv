package repo

import (
	"api/internal/config"
	"api/types"
	"context"
	"fmt"
	"time"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repo struct {
	storage *sqlx.DB
}

func New(config config.Config) *Repo {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.Username, config.Password, config.Database, config.Host, config.Port)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return nil
	}
	return &Repo{
		storage: db,
	}
}

func (r *Repo) CurrencyExists(ctx context.Context, to string) (bool, error) {
	query := `select exists(select from newtable where currency_to = :to)`
	_, err := r.storage.NamedExecContext(ctx, query, to)
	if err != nil {
		return true, nil
	}
	return true, nil
}

func (r *Repo) WriteRow(ctx context.Context, data types.DataPost) error {
	dataWrite := types.DatabaseFields{
		CurrencyFrom: data.CurrencyFrom,
		CurrencyTo:   data.CurrencyTo,
		Well:         0,
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}
	query := `INSERT INTO newtable (currency_from, currency_to, well, updated_at) 
        	VALUES(:currency_from, :currency_to, :well, :updated_at)`
	_, err := r.storage.NamedExecContext(ctx, query, dataWrite)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateWell(ctx context.Context, newData types.DataPut) error {
	query := `UPDATE newtable
			set well =:well
			where currency_to =:currency_to`
	_, err := r.storage.NamedExecContext(ctx, query, newData)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) UpdateRows(ctx context.Context, newData types.CurrencyLatest) error {
	query := `UPDATE newtable
			set well = :well,
			updated_at = :updated_at
			where currency_to =:currency_to`

	_, err := r.storage.NamedExecContext(ctx, query, newData.Data)
	return err
}

func (r *Repo) SelectRow(ctx context.Context, data types.DatabaseFields) (types.DatabaseFields, error) {
	query := `SELECT currency_from, currency_to, well, updated_at
			FROM newtable
			WHERE currency_from =:currency_from AND currency_to =:currency_to`
	res := types.DatabaseFields{}

	q, args, err := r.storage.BindNamed(query, data)
	if err != nil {
		return types.DatabaseFields{}, err
	}

	nq := r.storage.Rebind(q)
	if err := r.storage.GetContext(ctx, &res, nq, args...); err != nil {
		return res, err
	}
	return res, nil
}

func (r *Repo) SelectAll(ctx context.Context) ([]types.DatabaseFields, error) {
	query := `SELECT currency_from, currency_to, well, updated_at FROM newtable`
	var res []types.DatabaseFields
	if err := r.storage.GetContext(ctx, &res, query); err != nil {
		return res, err
	}
	return res, nil
}
