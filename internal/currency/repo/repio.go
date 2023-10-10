package repo

import (
	"api/internal/config"
	"context"
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Repo struct {
	storage *sqlx.DB
}

func New(config config.Config) (*Repo, error) {
	s := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable", config.Username, config.Password, config.Database, config.Host, config.Port)
	db, err := sqlx.Connect("postgres", s)
	if err != nil {
		return nil, err
	}
	return &Repo{
		storage: db,
	}, nil
}

func (r *Repo) CurrencyExists(ctx context.Context, to string) (ok bool, err error) {
	query := `select exists(select from newtable where currency_to = :to)`
	res, err := r.storage.NamedExecContext(ctx, query, to)
	fmt.Println(res)
	return true, nil
}
