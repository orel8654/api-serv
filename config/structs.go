package config

type ConfDB struct {
	Database string `yaml:"dbname"`
	Username string `yaml:"dbuser"`
	Password string `yaml:"dbpass"`
	Host     string `yaml:"dbhost"`
	Port     string `yaml:"dbport"`
}

type ConfAPI struct {
	Token string `yaml:"token"`
}

type CurrencyVal struct {
	RUB float64 `json:"RUB"`
	EUR float64 `json:"EUR"`
}

type CurrencyLatest struct {
	Data CurrencyVal `json:"rates"`
}

type DatabaseFields struct {
	CurrencyFrom string  `db:"currency_from"`
	CurrencyTo   string  `db:"currency_to"`
	Well         float64 `db:"well"`
	UpdatedAt    string  `db:"updated_at"`
}

type ResponseFields struct {
	CurrencyFrom string  `db:"currency_from"`
	CurrencyTo   string  `db:"currency_to"`
	Well         float64 `db:"well"`
}

type DataPost struct {
	CurrencyFrom string `json:"currency_from"`
	CurrencyTo   string `json:"currency_to"`
}

type DataPut struct {
	CurrencyFrom string  `json:"currency_from"`
	CurrencyTo   string  `json:"currency_to"`
	Well         float64 `json:"well"`
}
