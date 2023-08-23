package database

import (
	"database/sql"
	"log"
)

type EmiroSQLLite struct {
	Con *sql.DB
}

func (e *EmiroSQLLite) SearchQuery(query string, count int32, all bool) (*sql.Rows, error) {
	stmt, err := e.Con.Prepare(`SELECT name, description, command, language, path, script, interactive, os, author FROM Queries WHERE name LIKE ? OR description LIKE ?`)

	if err != nil {
		log.Fatal("prepared", err)
	}

	rows, err := stmt.Query("%"+query+"%", "%"+query+"%")

	if err != nil {
		log.Fatal("queried", err)
	}

	return rows, err

}

func (e *EmiroSQLLite) CreateOrConnect() error {
	var err error
	e.Con, err = sql.Open("sqlite3", "emiro.sqlite")
	if err != nil {
		log.Panic(err)
		return err
	}
	_, err = e.Con.Exec(sqlStatement)
	if err != nil {
		log.Panic(err)
		return err
	}

	return nil
}
