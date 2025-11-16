package database

import (
	"context"
	"go-fullstack/config"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/rs/zerolog"
)

func CreateDBPool(conf *config.DBConfig, logger *zerolog.Logger) *pgxpool.Pool {
	dbpool, err := pgxpool.New(context.Background(), conf.Url)

	if err != nil {
		logger.Error().Msg("Не удалось подключиться к базе данных " + err.Error())
		panic(err)
	}

	logger.Info().Msg("Подключение к БД прошло успешно")
	return dbpool
}
