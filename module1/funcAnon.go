package module1

import "fmt"

type FuncAnonymous func(string) bool

func faga(data string, tanda FuncAnonymous) {
	fmt.Println(tanda(data))
}

func FuncAnon()  {
	funcAnonymous := func (data string)bool  {
		return data == "naufal"
	}
	faga("nau",funcAnonymous)

	faga("naufal",func(s string) bool {
		return len(s) > 3
	})

	// this closure function
	result:= closureFunc(5)
	dar := result(4)

	fmt.Println("ini hasil closure",dar)
	
}

func closureFunc(x int) func(int)int{
	return func (y int) int {
		return x+y
	}
}