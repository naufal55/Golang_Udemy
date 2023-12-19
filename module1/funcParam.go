package module1

import "fmt"

// 2 parameter yang pertama string dan yang kedua adalah func dengan 
// parameter string dan value string
func param(data string, temp func(string) string) {
	
	fmt.Println("ini data",temp(data))
}

// optional
func call(name string)string  {
	return "asli data "+ name
}

func param2(data string, temp2 func(string)string)  {
	filterName := temp2(data)

	fmt.Println("ini asli",filterName)
}

func FuncParam()  {
	param("naufal",func(s string) string {
		return "asli data " + s
	})

	param2("aufa", call)
}