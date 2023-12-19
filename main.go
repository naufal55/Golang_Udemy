package main

import (
	// luarmodul "github.com/naufal55/golang_udemy/luarModul"
	"fmt"

	"github.com/naufal55/golang_udemy/module1"
	"github.com/naufal55/golang_udemy/module2"

	//ingin mengimport file go dengan isian function init saja tanpa ada daalamnya
	// maka pakai blank identifier "_" yg berarti eksekusi ini tanpa func dalamnya
	_ "github.com/naufal55/golang_udemy/database"
)

func main() {
	SemuaModul1()
}

type data struct{
	Nama string
	Age int
	Address string
}

type person struct{
	bundleData *data
}

func SemuaModul1()  {

	// module1.HelloWorld()
	// module1.Variabel()
	// module1.ConvertDataType()
	// module1.TypeDeclaration()
	// module1.ArrayData()
	// module1.MapData()

	// a,b := luarmodul.MapStruct()
	// fmt.Println("ini hasil",a,b)
	// dataArray:= data{
	// 	Nama: "Naufal",
	// 	Age: 20,
	// 	Address: "Jakarta",
	// }
	// module1.TestVariadic(dataArray)

	// module1.FuncParam()
	// module1.FuncAnon()
	fmt.Println(module1.Recursive(5))
	// module1.TypeAssertion()
	// module1.DeferPanicAndRecover(true)
	// module1.StructMethod()
	// module1.PointerRef()
	// module1.InitFunc()
	// luarmodul.Action()

	// module1.DataError()


	//-----------------------module 2--------------------------

	module2.SortData()

}