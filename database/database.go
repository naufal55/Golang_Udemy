package database

import "fmt"

var connection string

func init() {
	connection = "database PGSQL connected"
	fmt.Println(connection)
}