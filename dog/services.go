package dog

import (
	"fmt"
	"infinity-dog/database"
	"infinity-dog/network"
	"os"
)

var servicesExceptions = []string{}
var servicesMessages = []string{}

func ServicesFromSql(sortString, service string) {
	serviceItems := database.ServicesByTotalBytes()
	for i, item := range serviceItems {
		fmt.Printf("%03d. %-60s %d\n", i+1, item.Name, item.TotalBytes)
	}
}

func ServiceDependencies() {
	//api/v1/service_dependencies?env=
	jsonString := network.DoGet("/api/v1/service_dependencies?env=" + os.Getenv("DOG_PROD"))
	fmt.Println(jsonString)
}
