package luarmodul

import (
	"fmt"
	"reflect"
)

type data struct {
	name string
	age  int
}

func MapStruct() (person data, v reflect.Value){ // deklarasi return func sekaligus declare variabel
	person = data{"naufal", 20}

	v = reflect.ValueOf(person) // menampilkan value dari data struct: {naufal 20}

	t := v.Type() // menampilkan data struct mana dia berasal: luarmodul.data

	fmt.Println("ini jumlah field", t.NumField()) // menampilkan jumlah field dalam struct : 2

	// field mengambil urutan/iterasi keberapa field yg diambil dari struct
	fmt.Println("data ke 1", v.Field(0))

	// field mengambil urutan/iterasi keberapa field yg diambil dari struct
	fmt.Println("data ke 2", v.Field(1))

	ky1 := t.Field(1)
	// klo diprint 		{age 	github.com/naufal55/golang_udemy/luarModul 	int  	16 		[1] 	false}
	// isinya ky1 => 	.Name 	.PkgPath									.type	.offset	.Index	.IsExported

	// sama aja bedanya ini index keberapa setelah di iterasikan
	fmt.Println("key ke 1", ky1.Index)

	// sama aja bedanya ini keynya
	fmt.Println("key ke 1", ky1.Name)

	return person, v
}
