package database

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var (
	DB           EmiroDatabase
	sqlStatement string = `
		CREATE TABLE IF NOT EXISTS Queries
		(
			name TEXT,
			description TEXT,
			command TEXT,
			language VARCHAR(20),
			path VARCHAR(255),
			script INTEGER(1),
			interactive INTEGER(1),
			os TEXT,
			author TEXT
		);
	`
)

type EmiroDatabase interface {
	CreateOrConnect() error
	SearchQuery(query string, count int32, all bool) (*sql.Rows, error)
}

func InitDatabase(typ string) EmiroDatabase {
	switch typ {
	case "sqlite":
		sqlInstance := EmiroSQLLite{}
		sqlInstance.CreateOrConnect()
		return &sqlInstance
	case "elasticsearch":
		sqlInstance := EmiroElasticsearch{}
		sqlInstance.CreateOrConnect()
	}

	return nil
}
