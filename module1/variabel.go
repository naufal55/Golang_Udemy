package module1

import "reflect"

func Variabel() {
	var nama string = "naufal"
	number := 3
	tipe := reflect.TypeOf(number)

	var isReturn bool = true

	println(nama, len(nama),nama[5], number, isReturn, tipe.String())
}