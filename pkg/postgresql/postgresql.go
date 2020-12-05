//go:generate go run -v github.com/go-bindata/go-bindata/v3/go-bindata -pkg postgresql -o bindata.go migration

package postgresql

import (
	"fmt"
	"log"
	"path"

	"github.com/diazharizky/scheduler/internal/utils"
	"github.com/golang-migrate/migrate/v4"
	bindata "github.com/golang-migrate/migrate/v4/source/go_bindata"
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

const (
	migrationDir = "migration"
)

// Open func
func (p *PGInstance) Open() {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", p.Host, p.Port, p.User, p.Password, p.Database)
	conn := sqlx.MustConnect("pgx", dsn)
	p.Conn = conn
}

// MigrateUp func
func (p *PGInstance) MigrateUp() error {
	ls, err := AssetDir(migrationDir)
	if err != nil {
		log.Fatal(err.Error())
	}

	i := bindata.Resource(ls, func(name string) ([]byte, error) {
		return Asset(path.Join(migrationDir, name))
	})

	d, err := bindata.WithInstance(i)
	if err != nil {
		log.Fatal(err.Error())
	}

	dsn := utils.GetDSN("postgres", p.User, p.Password, p.Host, p.Port, p.Database, false)
	m, err := migrate.NewWithSourceInstance("go-bindata", d, dsn)
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
	dsn := utils.GetDSN("postgres", p.User, p.Password, p.Host, p.Port, p.Database, false)
	m, err := migrate.New("file://internal/migrations/postgres", dsn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err = m.Down(); err != nil {
		log.Fatal(err.Error())
	}

	return err
}
