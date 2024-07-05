package sql

import (
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/stretchr/testify/assert"
)

func TestSQLDatabase_Connect(t *testing.T) {
	mockDB, _, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}
	assert.NotNil(t, sqlDB.DB)
}

func TestSQLDatabase_Close(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	mock.ExpectClose()

	err = sqlDB.Close()
	assert.NoError(t, err)
}

func TestSQLDatabase_Query(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	rows := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "John Doe")
	mock.ExpectQuery("SELECT id, name FROM users").WillReturnRows(rows)

	result, err := sqlDB.Query("SELECT id, name FROM users")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSQLDatabase_QueryRow(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	row := sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "John Doe").RowError(0, nil)
	mock.ExpectQuery("SELECT id, name FROM users WHERE id = ?").WithArgs(1).WillReturnRows(row)

	result := sqlDB.QueryRow("SELECT id, name FROM users WHERE id = ?", 1)
	assert.NotNil(t, result)
}

func TestSQLDatabase_Exec(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	mock.ExpectExec("INSERT INTO users").WithArgs("John Doe").WillReturnResult(sqlmock.NewResult(1, 1))

	result, err := sqlDB.Exec("INSERT INTO users (name) VALUES (?)", "John Doe")
	assert.NoError(t, err)
	assert.NotNil(t, result)
}

func TestSQLDatabase_Ping(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	mock.ExpectPing()

	err = sqlDB.Ping()
	assert.NoError(t, err)
}

func TestSQLDatabase_Insert(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	mock.ExpectExec("INSERT INTO users").WithArgs("John Doe").WillReturnResult(sqlmock.NewResult(1, 1))

	err = sqlDB.Insert("INSERT INTO users (name) VALUES (?)", "John Doe")
	assert.NoError(t, err)
}

func TestSQLDatabase_Update(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	mock.ExpectExec("UPDATE users SET name = \\? WHERE id = \\?").WithArgs("John Doe", 1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = sqlDB.Update("UPDATE users SET name = ? WHERE id = ?", "John Doe", 1)
	assert.NoError(t, err)
}

func TestSQLDatabase_Delete(t *testing.T) {
	mockDB, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer mockDB.Close()

	sqlxDB := NewDb(mockDB, "sqlmock")
	sqlDB := &SQLDatabase{DB: sqlxDB}

	mock.ExpectExec("DELETE FROM users WHERE id = \\?").WithArgs(1).WillReturnResult(sqlmock.NewResult(1, 1))

	err = sqlDB.Delete("DELETE FROM users WHERE id = ?", 1)
	assert.NoError(t, err)
}
