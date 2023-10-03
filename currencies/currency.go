package currencies

import "api/config"

type Currency struct {
	jwt config.ConfAPI
}

func NewCurrency(conf config.ConfAPI) *Currency {
	return &Currency{
		jwt: conf,
	}
}

func (c *Currency) GetNewCurrency() {

}
