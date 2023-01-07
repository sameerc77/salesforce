package main

import (
	"fmt"

	"github.com/sameerc77/salesforce/connection"
)

func main() {
	fmt.Println(connection.GetSalesforceConn())
}
