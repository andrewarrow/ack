package dog

import (
	"fmt"
	"infinity-dog/network"
	"time"
)

func Metrics() {
	jsonString := network.DoGet(fmt.Sprintf("/api/v1/metrics?from=%d", time.Now().Unix()-86400))
	fmt.Println(jsonString)
}
