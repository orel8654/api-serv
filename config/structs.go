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

type DatabaseFields struct {
	CurrencyFrom string  `pg:"currency_from"`
	CurrencyTo   string  `pg:"currency_to"`
	Well         float64 `pg:"well"`
	UpdatedAt    string  `pg:"updated_at"`
}

type ResponseFields struct {
	CurrencyFrom string  `pg:"currency_from"`
	CurrencyTo   string  `pg:"currency_to"`
	Well         float64 `pg:"well"`
}
