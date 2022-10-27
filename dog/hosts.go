package dog

import (
	"fmt"
	"infinity-dog/network"
)

func ActiveHosts() {
	jsonString := network.DoGet("/api/v1/hosts")
	fmt.Println(jsonString)
}
