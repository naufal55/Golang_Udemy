package module1

import "fmt"

func log()  {
	
	fmt.Println("aplikasi berakhir")
	mess := recover()
	// why in this code show error mess is <nil>?
	fmt.Println("pesan error :", mess)
}

func DeferPanicAndRecover(err bool) {

	defer log()

	if err {
		panic("ups Error")
	}

}