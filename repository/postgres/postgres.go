package postgres

import (
	"fmt"

	"github.com/jmoiron/sqlx"

	"github.com/madxiii/tsarka_task/configs"
)

func UserConn(cfg configs.Postgres) (*User, error) {
	dsn := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.Username, cfg.DBname, cfg.Password, cfg.SSLMode,
	)

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		return nil, err
	}

	if err = db.Ping(); err != nil {
		return nil, err
	}

	return &User{db: db}, nil
}
