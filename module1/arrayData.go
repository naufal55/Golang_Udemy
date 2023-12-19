package module1

import (
	"fmt"
	"reflect"
)

func ArrayData() {
	
	// bedakan ARRAY DAN SLICE !!!!
	arr := [...]int{3,2,4} // array
	slc := []int{1,2,3}    // slice
	fmt.Println(reflect.TypeOf(arr), reflect.TypeOf(slc))

	// cara menggunakan copy() array /  slice
	var gag = make([]int,len(slc),cap(slc))
	copy(gag, slc)
	fmt.Println("after copy arr: ", gag)
	fmt.Println("after copy slc: ", slc)

	//------------------------------------------------------------------------------
	// 1 . array tidak bisa di append sendiri dan harus ditampung yang baru
	// sedangkan slice bisa
	a:=[]int{1,2,3}
	a = append(a,4)
	fmt.Println(a)

	b:=[...]int{1,2,3}
	//b = append(b, 4) // Error : first argument to append must be a slice;
	fmt.Println(b)
	
	
	//--------------------------------------------------------------------------------
	//2. perbedaan array dan slice	
	//create array
	// ketika di append array kedua, slice pertama tetap tidak berubah
	jalla := [...]string{"ahjaksld", "blaksd"}

	vava := append(jalla[:], "data baru")
	vava[0] = "berubah"
	fmt.Println(jalla, vava)

	// create slice with length 2 dan capacity maks sampai 5 buat di append
	// ketika di append slice kedua maka slice pertama akan berubah
	var newSlice []string = make([]string, 2,5)

	newSlice[0] = "haha"
	newSlice[1] = "ada"

	glo := append(newSlice, "fafafaf")
	glo[0] = "rubah"

	fmt.Println(glo, newSlice)

	//--------------------------------------------------------------------------------
	//3. slice bisa dibuat dengan dengan array[:] 
	//atau bisa dengan langsng diawal []int{blabla}
	//
	// sedangkan aray harus fixed dengan contoh : "[2]int{blbabla}" 
	//atau [...]int{blablabla}

	c:= [...]int{1,2,3,4,5,6,7} // create array langsung
	cc := c[:] //create slice dengan slice

	fmt.Println(c, cc)

	switch {
	case len(c)>2 :
		fmt.Println("data pas")
	default:
		fmt.Println("data tidak pas")
	}

	for _, data := range c {
		if data == 3{
			continue
		}
		fmt.Println(data)
		if data == 6 {
			break;
		}
	}
	// 4. array hanya bisa menggunakan len()
	//sedangkan slice bisa menggunakan len() dan cap() for capacity
	fmt.Println("len of c: ", len(c))
	fmt.Println("capacity of c: ", cap(c))

}