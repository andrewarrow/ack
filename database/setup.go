package database

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func OpenTheDB() *sql.DB {
	db, err := sql.Open("sqlite3", "sqlite.db")
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return db
}

func CreateSchema() {
	db := OpenTheDB()
	defer db.Close()

	// explain query plan
	// sqlite version 3.39.3
	sqlStmt := `
CREATE TABLE services (id text, 
                       name text, 
											 msg text, 
											 message text, 
											 exception text, 
											 logged_at datetime);
CREATE INDEX IF NOT EXISTS index1 ON services (name);
CREATE INDEX IF NOT EXISTS index2 ON services (name,logged_at);
CREATE UNIQUE INDEX IF NOT EXISTS index3 ON services (id);

CREATE TABLE service_meta (name text, 
                       total_exceptions integer,
                       total_bytes integer);
CREATE UNIQUE INDEX IF NOT EXISTS index4 ON service_meta (name);
	`
	_, err := db.Exec(sqlStmt)
	if err != nil {
		fmt.Printf("%q\n", err)
		return
	}
}
