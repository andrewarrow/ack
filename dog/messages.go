package dog

import (
	"fmt"
	"infinity-dog/database"
	"time"
)

func Messages(serviceName string) {
	items := database.MessagesFromService(serviceName)
	utcNow := time.Now().In(utc).Unix()
	for _, item := range items {
		delta := float64(utcNow-item.LoggedAt) / 3600.0
		fmt.Printf("%.2f %d %s\n", delta, len(item.Both), item.BothTruncated(0))
	}
}
