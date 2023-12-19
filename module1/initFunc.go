package module1

import "fmt"

var connection string
// init akan di eksekusi di file go untuk pertama kali sebelum function lain seperti 
//constructor

func init() {
	connection = "MYSQL"
}
func InitFunc() string {
	fmt.Println(connection)
	return connection
}