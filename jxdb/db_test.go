package jxdb

import (
	"log"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
	"github.com/upperxcode/go_jxdb/pkg/sql"
)

func TestDB_ConnectAndClose(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)

	sqlxDB := sql.NewDb(mockDB, "sqlmock")
	sqlDB := &sql.SQLDatabase{DB: sqlxDB}

	Init = func(driverName string, dataSourceName string) (sql.Database, error) {
		return sqlDB, nil
	}

	driver := Postgres
	host := "172.18.0.2"
	user := "postgres"
	dbname := "upper"
	password := "postgres"
	port := 5432

	database, err := InitInstance(driver, host, user, dbname, password, port)
	if err != nil {
		log.Fatal(err)
	}
	assert.NoError(t, err)
	assert.NotNil(t, database.Conn)

	mock.ExpectClose()
	err = database.Close()
	assert.NoError(t, err)
}

func TestDB_buildConnStr(t *testing.T) {
	tests := []struct {
		name     string
		dbConfig DB
		want     string
		wantErr  bool
	}{
		{
			name: "Postgres connection string",
			dbConfig: DB{
				Driver:   Postgres,
				Host:     "localhost",
				User:     "user",
				Dbname:   "dbname",
				Password: "password",
			},
			want:    "host=localhost user=user dbname=dbname password=password sslmode=disable",
			wantErr: false,
		},
		{
			name: "Unsupported driver",
			dbConfig: DB{
				Driver: Driver(999),
			},
			want:    "",
			wantErr: true,
		},
		{
			name: "SQLite connection string",
			dbConfig: DB{
				Driver: SqLite,
				Dbname: "test.db",
			},
			want:    "test.db",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.dbConfig.buildConnStr()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
