package module2

import (
	"encoding/base64"
	"fmt"
	"sort"
)

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func (a UserSlice) Len() int           { return len(a) }
func (a UserSlice) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a UserSlice) Less(i, j int) bool { return a[i].Age < a[j].Age }

func SortData() {
	users := []User{
		{"Tom", 20},
		{"Jerry", 18},
		{"Alice", 30},
	}
	 
	dt := UserSlice(users)
	fmt.Println(dt)
	sort.Sort(dt) // manggil constructor menegnai ini
	fmt.Println([]User(users))
	// fmt.Println()
	da := dt.Len()
	fmt.Println(da)


	value := "muhammad naufal"

	encodeString := base64.StdEncoding

	encoded := encodeString.EncodeToString([]byte(value))

	fmt.Println(encoded)

	decoded, err := encodeString.DecodeString(encoded)

	if err != nil {
		fmt.Println("error",err.Error())
	}else{
		fmt.Println(string(decoded))
	}
}