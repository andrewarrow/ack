package database

import (
	"fmt"
	"strings"
)

type Message struct {
	Both            string
	ExceptionLength int
	LoggedAt        int64
	Id              string
}

func MessagesFromService(service string) []Message {
	items := []Message{}
	s := fmt.Sprintf(`select id, msg, message, length(exception) as exception_length, unixepoch(logged_at) as ts from services where name='%s' order by logged_at desc limit 300`, service)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	for rows.Next() {
		var msg string
		var messageString string
		var exceptionLength int
		var loggedAt int64
		var id string
		rows.Scan(&id, &msg, &messageString, &exceptionLength, &loggedAt)
		message := Message{}
		message.Both = msg + messageString
		message.ExceptionLength = exceptionLength
		message.LoggedAt = loggedAt
		message.Id = id
		items = append(items, message)
	}

	return items
}

func (m *Message) BothTruncated(offset int) string {
	if len(m.Both) > 90 {
		return strings.ReplaceAll(m.Both[0:90], "\n", " ")
	}
	return strings.ReplaceAll(m.Both, "\n", " ")
}
