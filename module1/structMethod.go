package module1

import "fmt"

type data struct {
	Nama, Alamat string
	umur         int
}

func (d data) getData() {
	fmt.Println("nama saya", d.Nama, " alamat saya : ",d.Alamat, " ,dan saya berumur ", d.umur )
}




func StructMethod() {
	joko := data{
		Nama:  "Joko",
		Alamat:"Bandung",
		umur:  20,
	}

	// atau bisa seperti ini
	budi := data{"budi","bandung",23}

	joko.getData()
	budi.getData()
}