package repo

import "context"

type Repo struct {
}

func New() *Repo {
	return &Repo{}
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
