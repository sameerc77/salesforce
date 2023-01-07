package main

import (
	"fmt"

	"github.com/sameerc77/salesforce/connection"
	"github.com/sameerc77/salesforce/db"
	"github.com/sameerc77/salesforce/logger"
)

func main() {
	fmt.Println(connection.GetSalesforceConn(), db.GetConnection(), logger.Getlog())
}
