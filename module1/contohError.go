package module1

import (
	"errors"
	"fmt"
)

func errorExample(data interface{}) (any, error) {
	
	kl,err := data.(string)
	// if data type not string return error
	if _,ok := data.(string); ok {
		// iforget if data.(string) must be 2 return value value and bool (can be converted or not)
		return nil, errors.New("data type not string")
	}
	fmt.Println("data",kl,err)
	
	return data, nil
}

func DataError()  {
	data, err := errorExample(3)
	if err != nil {
		fmt.Println("ini error", err)
	} else {
		fmt.Println("ini data", data)
	}
}