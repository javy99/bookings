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

// testDBRepo
type testDBRepo struct {
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

// NewTestingRepo
func NewTestingRepo(a *config.AppConfig) repository.DatabaseRepo {
	return &testDBRepo{
		App: a,
	}
}
