package currencies

import (
	"api/config"
	"encoding/json"
	"net/http"
)

const UrlsEx = "openexchangerates.org/api/latest.json?base=USD&app_id="

type Currency struct {
	jwt config.ConfAPI
}

func NewCurrency(conf config.ConfAPI) *Currency {
	return &Currency{
		jwt: conf,
	}
}

func (c *Currency) GetNewCurrency() (config.CurrencyLatest, error) {
	var resultData config.CurrencyLatest
	response, err := http.Get(UrlsEx + c.jwt.Token)
	if err != nil {
		return resultData, err
	}
	defer response.Body.Close()
	if err = json.NewDecoder(response.Body).Decode(&resultData); err != nil {
		return resultData, err
	}
	return resultData, nil
}
