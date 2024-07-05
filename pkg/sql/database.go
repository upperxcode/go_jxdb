package sql

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq" // Importar o driver do PostgreSQL
)

type Database interface {
	Connect(driverName string, dataSourceName string) error
	Close() error
	Query(query string, args ...interface{}) (*sqlx.Rows, error)
	QueryRow(query string, args ...interface{}) *sqlx.Row
	Select(dest interface{}, query string, args ...interface{}) error
	Get(dest interface{}, query string, args ...interface{}) error
	Exec(query string, args ...interface{}) (sql.Result, error)
	Ping() error
	List(dest interface{}, where string, args ...interface{}) error
	Insert(query string, args ...interface{}) error
	Update(query string, args ...interface{}) error
	Delete(query string, id int) error
}

// Implementação concreta da interface Database usando sqlx.DB
type SQLDatabase struct {
	DB *sqlx.DB
}

func (d *SQLDatabase) Connect(driverName string, dataSourceName string) error {
	var err error
	d.DB, err = sqlx.Connect(driverName, dataSourceName)
	return err
}

func (d *SQLDatabase) Close() error {
	return d.DB.Close()
}

func (d *SQLDatabase) Query(query string, args ...interface{}) (*sqlx.Rows, error) {
	return d.DB.Queryx(query, args...)
}

func (d *SQLDatabase) QueryRow(query string, args ...interface{}) *sqlx.Row {
	return d.DB.QueryRowx(query, args...)
}

func (d *SQLDatabase) Select(dest interface{}, query string, args ...interface{}) error {
	return d.DB.Select(dest, query, args...)
}

func (d *SQLDatabase) Get(dest interface{}, query string, args ...interface{}) error {
	println(query)
	return d.DB.Get(dest, query, args...)
}

func (d *SQLDatabase) Exec(query string, args ...interface{}) (sql.Result, error) {
	return d.DB.Exec(query, args...)
}

func (d *SQLDatabase) Ping() error {
	return d.DB.Ping()
}

func (d *SQLDatabase) List(dest interface{}, where string, args ...interface{}) error {
	query := "SELECT * FROM " + where
	return d.Select(dest, query, args...)
}

func (d *SQLDatabase) Insert(query string, args ...interface{}) error {
	_, err := d.Exec(query, args...)
	return err
}

func (d *SQLDatabase) Update(query string, args ...interface{}) error {
	_, err := d.Exec(query, args...)
	return err
}

func (d *SQLDatabase) Delete(query string, id int) error {
	_, err := d.Exec(query, id)
	return err
}

func NewDb(db *sql.DB, driverName string) *sqlx.DB {

	return sqlx.NewDb(db, "sqlmock")
}
