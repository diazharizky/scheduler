package postgres

import (
	"fmt"
	"log"

	// Use pgx as Postgres's DB driver
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
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
func (p *PGInstance) Open() error {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Database)
	conn, err := sqlx.Connect("pgx", dsn)
	if err != nil {
		log.Fatalln(err)
	}
	p.Conn = conn

	return err
}

// CreateSchedule func
func (p *PGInstance) CreateSchedule() (err error) {
	return err
}
