package main

import (
	"fmt"

	"github.com/sameerc77/salesforce/connection"
	"github.com/sameerc77/salesforce/db"
)

func main() {
	fmt.Println(connection.GetSalesforceConn(), db.GetConnection())
}
