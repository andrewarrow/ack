package database

import "strconv"

type Service struct {
	Name       string
	TotalBytes int64
}

func ServicesByTotalBytes() []Service {
	items := []Service{}
	s := `select name, total_bytes from service_meta order by total_bytes desc`
	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	for rows.Next() {
		var name string
		var totalBytes string
		rows.Scan(&name, &totalBytes)
		service := Service{}
		service.Name = name
		service.TotalBytes, _ = strconv.ParseInt(totalBytes, 10, 64)
		items = append(items, service)
	}

	return items
}
