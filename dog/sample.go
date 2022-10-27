package dog

import (
	"os"
	"strconv"
)

func Sample(hours string) {
	hoursAsInt, _ := strconv.Atoi(hours)
	Logs(hoursAsInt, os.Getenv("DOG_BASE"))
}
