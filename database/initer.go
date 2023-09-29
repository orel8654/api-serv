package database

import (
	"api/config"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	"net"
)

func Connect(conf *config.ConfDB) (*pg.DB, error) {
	db := pg.Connect(&pg.Options{
		Addr:     net.JoinHostPort(conf.Host, conf.Port),
		User:     conf.Username,
		Password: conf.Password,
		Database: conf.Database,
	})
	if err := CreateSchema(db); err != nil {
		return nil, err
	}
	return db, nil
}

func CreateSchema(db *pg.DB) error {
	models := []interface{}{
		(*config.DatabaseFields)(nil),
	}
	for _, model := range models {
		err := db.Model(model).CreateTable(&orm.CreateTableOptions{
			Temp: true,
			})
		if err != nil {
			return err
		}
	}
	return nil
}