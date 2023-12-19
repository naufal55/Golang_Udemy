package helper

import "fmt"

func ChecksExist(value any) (any, error) {
	if value == nil {
		return "",fmt.Errorf("error data null")
	}
	return value, nil
}