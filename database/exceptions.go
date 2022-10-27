package database

import (
	"fmt"
)

type Exception struct {
	Text     string
	LoggedAt int64
}

func ExceptionsFromService(service string) []Exception {
	items := []Exception{}
	s := fmt.Sprintf(`select exception, unixepoch(logged_at) as ts from services where name='%s' order by logged_at desc limit 60`, service)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	for rows.Next() {
		var exceptionString string
		var loggedAt int64
		rows.Scan(&exceptionString, &loggedAt)
		e := Exception{}
		e.Text = exceptionString
		e.LoggedAt = loggedAt
		items = append(items, e)
	}

	return items
}

func ExceptionById(id string) Exception {
	s := fmt.Sprintf(`select exception from services where id='%s'`, id)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	rows.Next()

	var exceptionString string
	rows.Scan(&exceptionString)

	e := Exception{}
	e.Text = exceptionString

	return e
}
