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
	res, err := r.storage.NamedExecContext(ctx, query, to)
	if err != nil {
		return true, nil
	}
	fmt.Println(res)
	return true, nil
}
