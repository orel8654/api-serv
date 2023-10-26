package postgres

import (
	"context"
	"time"

	"github.com/jmoiron/sqlx"

	"api/internal/types"
)

type Repo struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) *Repo {
	return &Repo{
		db: db,
	}
}

// func (r *Repo) CurrencyExists(ctx context.Context, from, to string) (ok bool, err error) {
// 	const query = `
// select exists (
// 	select from currency where currency_from=:from and currency_to = :to
// )
// `
// 	return false, nil
// }

func (r *Repo) UpdateWell(ctx context.Context, newData types.DataPut) error {
	query := `
UPDATE newtable
	set well =$1
where currency_to =$2
`
	_, err := r.db.ExecContext(ctx, query, newData.Well, newData.CurrencyTo)
	return err
}

func (r *Repo) WriteRow(ctx context.Context, data types.DataPost) error {
	dataWrite := types.DatabaseFields{
		CurrencyFrom: data.CurrencyFrom,
		CurrencyTo:   data.CurrencyTo,
		Well:         0,
		UpdatedAt:    time.Now().Format("2006-01-02 15:04:05"),
	}

	query := `
INSERT INTO newtable (currency_from, currency_to, well, updated_at)
VALUES(:currency_from, :currency_to, :well, :updated_at)
`
	_, err := r.db.NamedExecContext(ctx, query, dataWrite)
	if err != nil {
		return err
	}
	return nil
}

// func (r *Repo) UpdateRows(ctx context.Context, newData types.CurrencyLatest) error {
// 	query := `
// UPDATE newtable
// 	SET well = :well,
// 		updated_at = :updated_at
// WHERE currency_to =:currency_to
// `

// 	_, err := r.db.NamedExecContext(ctx, query, newData.Data)
// 	return err
// }

// func (r *Repo) SelectRow(ctx context.Context, data types.DatabaseFields) (types.DatabaseFields, error) {
// 	query := `
// SELECT currency_from, currency_to, well, updated_at
// FROM newtable
// WHERE currency_from =:currency_from AND currency_to =:currency_to
// `
// 	res := types.DatabaseFields{}

// 	q, args, err := r.db.BindNamed(query, data)
// 	if err != nil {
// 		return types.DatabaseFields{}, err
// 	}

// 	nq := r.db.Rebind(q)

// 	if err := r.db.Get(&res, nq, args...); err != nil {
// 		return res, err
// 	}
// 	return res, nil
// }

func (r *Repo) SelectAll(ctx context.Context) ([]types.DatabaseFields, error) {
	query := `
SELECT currency_from, currency_to, well, updated_at
FROM newtable
`
	var res []types.DatabaseFields
	if err := r.db.SelectContext(ctx, &res, query); err != nil {
		return res, err
	}
	return res, nil
}
