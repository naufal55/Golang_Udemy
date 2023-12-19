// type assertion adalah perubahan tipe data yang pastinya adalah
// any atau interface kosong menjadi sebuah tipe data yg lain => string/int/dll
// apabila data yg dirubah tidak sesuai maka dia akan return panic

package module1

import "fmt"

func TypeAssertion()  {
	valueData := value().(string)

	// v:=value().(int) // akan keluar panic dari sini 
	// interface conversion: interface {} is string, not int   
	// fmt.Println(v)

	// maka dari itu solusinya adalah menggunakan switch case 
	// dan rubah type dangan (type) di dalam switch 
	switch value().(type) {
	case string:
		fmt.Println("nilai string",value().(string))
	case int:
		fmt.Println("nilai int",value().(int))
	case bool:
		fmt.Println("nilai bool",value().(bool))
	default:
		fmt.Println("Tidak terdefinisi")
	}
	

	fmt.Println(valueData)
}

func value() interface{} {
	return "bool"
}