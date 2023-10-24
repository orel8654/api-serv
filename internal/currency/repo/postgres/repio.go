package postgres

import (
	"context"

	"github.com/jmoiron/sqlx"

	"api/types"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

func (r *Repo) CurrencyExists(ctx context.Context, from, to string) (ok bool, err error) {
	const query = `
select exists (
	select from currency where currency_from=:from and currency_to = :to
)
`
	// Get Some from query
	return false, nil
}

func (r *Repo) UpdateWell(ctx context.Context, newData types.DataPut) error {
	query := `
UPDATE newtable
	set well =$1
where currency_to =$2
`
	_, err := r.db.ExecContext(ctx, query, newData.Well, newData.CurrencyTo)
	return err
}
