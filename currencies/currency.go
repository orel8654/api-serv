package currencies

import (
	"api/internal/types"
	"encoding/json"
	"net/http"
)

const UrlsEx = "https://openexchangerates.org/api/latest.json?base=USD&app_id="

type Currency struct {
	jwt types.ConfAPI
}

func NewCurrency(conf types.ConfAPI) *Currency {
	return &Currency{
		jwt: conf,
	}
}

func (c *Currency) GetNewCurrency() (types.CurrencyLatest, error) {
	var resultData types.CurrencyLatest
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
