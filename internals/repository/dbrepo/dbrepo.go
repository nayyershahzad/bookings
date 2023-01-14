package dbrepo

import (
	"database/sql"

	"github.com/nayyershahzad/bookings/internals/config"
)

type postgresDbRepo struct {
	App *config.Appconfig
	DB  *sql.DB
}

func NewPostgresRepo(conn *sql.DB, a *config.Appconfig) repository.DatabaseRepo
