package module1

import (
	"reflect"
	"strconv"
)

func ConvertDataType() {
	const (
		name    = "naufal"
		age     = "20"
		address = "jakarta"
		no_address = 10
	)

	//pick one char in string
	var char1 = string(name[0])

	var num,_ = strconv.Atoi(age)

	var noAddress = strconv.Itoa(no_address)

	println(char1, num, noAddress, reflect.TypeOf(num).String())

	
}