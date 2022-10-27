package dog

import (
	"fmt"
	"infinity-dog/database"
	"time"
)

func Exceptions(serviceName string) {
	items := database.ExceptionsFromService(serviceName)
	utcNow := time.Now().In(utc).Unix()
	for _, item := range items {
		delta := float64(utcNow-item.LoggedAt) / 3600.0
		fmt.Printf("%.2f %d\n", delta, len(item.Text))
		fmt.Println(item.Text)
	}
}
