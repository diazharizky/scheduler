package postgresql

import (
	"fmt"
	"log"

	"github.com/diazharizky/scheduler/internal/utils"
	"github.com/diazharizky/scheduler/pkg/server"
	"github.com/golang-migrate/migrate/v4"
	"github.com/jmoiron/sqlx"

	// Below are the required drivers
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/jackc/pgx/stdlib"
)

// PGInstance struct
type PGInstance struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string

	Conn *sqlx.DB
}

// Open func
func (p *PGInstance) Open() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Database)
	conn := sqlx.MustConnect("pgx", dsn)
	p.Conn = conn
}

// MigrateUp func
func (p *PGInstance) MigrateUp() error {
	dsn := utils.GetDSN("postgres", server.Config.GetString("POSTGRES_USER"), server.Config.GetString("POSTGRES_PASSWORD"), server.Config.GetString("POSTGRES_HOST"), server.Config.GetInt("POSTGRES_PORT"), server.Config.GetString("POSTGRES_DATABASE"), false)
	m, err := migrate.New("file://internal/migrations/postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = m.Up(); err != nil {
		log.Fatal(err.Error())
	}

	return err
}

// MigrateDown func
func (p *PGInstance) MigrateDown() error {
	dsn := utils.GetDSN("postgres", server.Config.GetString("POSTGRES_USER"), server.Config.GetString("POSTGRES_PASSWORD"), server.Config.GetString("POSTGRES_HOST"), server.Config.GetInt("POSTGRES_PORT"), server.Config.GetString("POSTGRES_DATABASE"), false)
	m, err := migrate.New("file://internal/migrations/postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = m.Down(); err != nil {
		log.Fatal(err.Error())
	}

	return err
}
