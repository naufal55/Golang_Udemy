package module1

import (
	"fmt"

	"github.com/naufal55/golang_udemy/helper"
)

func MapData()any{
	person := map[any]any{
		1:    3,
		"fa": "fu",
	}

	fmt.Println(person[1],person["fa"])

	// atau bisa menggunakan seprti ini

	var person1 map[any]any = map[any]any{}

	person1[1] = 3456
	person1["kk"] = "hghghjk"

	fmt.Println(person1)

	//untuk mendeklarasikan map kosong bisa menggunakan ini
	var person2 map[string]string = map[string]string{}
	person2["nama"] = "anjay"

	// atau dengan make
	person3 := make(map[string]int)
	// person3["nama"] = 9876
	le := person3["nama"]
	fmt.Println(le)

	value,err := helper.ChecksExist(nil)
	if err!=nil {
		fmt.Println(err)		
		return ""
	}
	return value

}
	


