package jxdb

import (
	"testing"
)

func TestDriverString(t *testing.T) {
	tests := []struct {
		driver   Driver
		expected string
	}{
		{Postgres, "postgres"},
		{MySql, "MySql"},
		{SqLite, "sqlite3"},
		{Oracle, "Oracle"},
		{Driver(999), "Unsupported DB Driver"}, // Teste para um valor desconhecido
	}

	for _, test := range tests {
		result := test.driver.String()
		if result != test.expected {
			t.Errorf("String() for driver %v: expected %s, got %s", test.driver, test.expected, result)
		}
	}
}

func TestDriverConnectionFormat(t *testing.T) {
	tests := []struct {
		driver   Driver
		expected string
	}{
		{Postgres, "host=%s user=%s dbname=%s password=%s sslmode=disable"},
		{MySql, "MySql"},
		{SqLite, "%s"},
		{Oracle, "Oracle"},
		{Driver(999), "unknown"}, // Teste para um valor desconhecido
	}

	for _, test := range tests {
		result := test.driver.ConnectionFormat()
		if result != test.expected {
			t.Errorf("ConnectionFormat() for driver %v: expected %s, got %s", test.driver, test.expected, result)
		}
	}
}
