package module1

import (
	"fmt"
	"reflect"
)

func TestVariadic(data any) {

	dataRaw := reflect.ValueOf(data)

	datar := []interface{}{}
	fmt.Println(data)

	// //this destruct struct dataArray to variabel datar
	
	for i := 0; i < dataRaw.NumField(); i++ {
		a:=dataRaw.Field(i).Interface()
		datar = append(datar,a)
	}
	result := variadic

	result(datar...)
}

func variadic(data ...any)  {

	for _, dataDiri := range data {
		fmt.Println(dataDiri)
	}
}