package dog

import (
	"fmt"
	"infinity-dog/network"
)

func Billing() {
	// auth needs `usage_read` scope
	//jsonString := network.DoGet("/api/v1/usage/billable-summary?month=2022-09")
	//jsonString := network.DoGet("/api/v2/usage/estimated_cost?view=sub-org&start_month=2021-07")
	jsonString := network.DoGet("/api/v2/usage/cost_by_org?start_month=2022-09")
	fmt.Println(jsonString)

}
