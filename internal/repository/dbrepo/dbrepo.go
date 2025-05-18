package dbrepo

import (
	"database/sql"

	"github.com/javy99/bookings/internal/config"
	"github.com/javy99/bookings/internal/repository"
)

// postgresDBRepo
type postgresDBRepo struct {
	App *config.AppConfig
	DB  *sql.DB
}

// NewPostgresRepo
func NewPostgresRepo(conn *sql.DB, a *config.AppConfig) repository.DatabaseRepo {
	return &postgresDBRepo{
		App: a,
		DB:  conn,
	}
}
