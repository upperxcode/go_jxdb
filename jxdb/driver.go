package jxdb

type Driver int

const (
	Postgres Driver = iota
	MySql
	SqLite
	Oracle
)

func (d Driver) String() string {
	switch d {
	case Postgres:
		return "postgres"
	case MySql:
		return "MySql"
	case SqLite:
		return "sqlite3"
	case Oracle:
		return "Oracle"
	default:
		return "Unsupported DB Driver"
	}

}

func (d Driver) ConnectionFormat() string {
	switch d {
	case Postgres:
		return "host=%s user=%s dbname=%s password=%s sslmode=disable"
	case MySql:
		return "MySql"
	case SqLite:
		return "%s"
	case Oracle:
		return "Oracle"
	default:
		return "unknown"
	}
}
