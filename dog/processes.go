package dog

import (
	"fmt"
	"infinity-dog/network"
)

func Processes() {
	jsonString := network.DoGet("/api/v2/processes")
	fmt.Println(jsonString)
}
