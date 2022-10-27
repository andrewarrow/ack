package database

import (
	"fmt"
	"time"
)

var utc, _ = time.LoadLocation("UTC")

func TotalRows() int64 {
	s := fmt.Sprintf(`select count(1) from services`)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	rows.Next()

	var total int64
	rows.Scan(&total)

	return total
}

func MinMaxDates() (float64, float64) {
	s := fmt.Sprintf(`select unixepoch(min(logged_at)), unixepoch(max(logged_at)) from services`)

	db := OpenTheDB()
	defer db.Close()

	rows, _ := db.Query(s)
	defer rows.Close()
	rows.Next()

	var min int64
	var max int64
	rows.Scan(&min, &max)

	utcNow := time.Now().In(utc).Unix()
	minDelta := float64(utcNow-min) / 86400.0
	maxDelta := float64(utcNow-max) / 86400.0

	return minDelta, maxDelta
}
