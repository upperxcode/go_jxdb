package jxdb

import (
	"fmt"
	"sync"

	"github.com/upperxcode/go_jxdb/pkg/sql"
)

type DB struct {
	Driver   Driver
	Host     string
	User     string
	Dbname   string
	Password string
	Port     int
	Conn     *Conn
}

type Conn struct {
	DB sql.Database
}

var (
	instance *DB
	//once     sync.Once
	initOnce sync.Once
	initErr  error
)

func InitInstance(driver Driver, host, user, dbname, password string, port int) (*DB, error) {
	initOnce.Do(func() {
		instance = &DB{
			Driver:   driver,
			Host:     host,
			User:     user,
			Dbname:   dbname,
			Password: password,
			Port:     port,
		}
		initErr = instance.connect()
	})
	return instance, initErr
}

// GetInstance retorna a instância única do DB.
func GetInstance() (*DB, error) {
	if instance == nil {
		return nil, fmt.Errorf("DB instance is not initialized. Call InitInstance first. %s", initErr)
	}
	return instance, nil
}

// Connect cria uma conexão com o banco de dados baseado no driver especificado.
func (d *DB) connect() error {
	connStr, err := d.buildConnStr()
	if err != nil {
		return err
	}

	db, err := Init(d.Driver.String(), connStr)
	if err != nil {
		return err
	}

	d.Conn = &Conn{DB: db}
	return nil
}

// Close fecha a conexão com o banco de dados.
func (d *DB) Close() error {
	if d.Conn != nil {
		return d.Conn.DB.Close()
	}
	return nil
}

// buildConnStr constrói a string de conexão baseada no driver.
func (d *DB) buildConnStr() (string, error) {
	switch d.Driver {
	case Postgres:
		return fmt.Sprintf("host=%s user=%s dbname=%s password=%s sslmode=disable", d.Host, d.User, d.Dbname, d.Password), nil
	case SqLite, MySql, Oracle:
		return d.Dbname, nil
	default:
		return "", fmt.Errorf("driver não suportado: %s", d.Driver)
	}
}

// Init inicializa a conexão com o banco de dados.
var Init = func(driverName string, dataSourceName string) (sql.Database, error) {
	db := &sql.SQLDatabase{}
	err := db.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}
	return db, nil
}
