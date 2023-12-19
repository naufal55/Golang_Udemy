package module1

import "fmt"

type dataSekian struct {
	name, age, address string
}

func PointerRef() {
	person1 := dataSekian{"naufal", "30", "jakarta"}
	//pakai pointer & agar reference ke pointer1
	person2 := &person1

	// mengacu ke adress 1 dan tidak bisa diisi langsung banyak field (harus 1 1)
	person2.name = "kololol"
	fmt.Println(person1)
	fmt.Println(person2)

	// ini sama aja seperti ini
	// |------------------|
	// |    dataSekian    |
	// |------------------|
	//	 |		  		
	//	 |		  		
	//person1----( & )-----person2
	// 0x001				0x001

	//----------------------------------------------------------------------------------

	// namun apabila dia membuat object/datasekian yg baru dengan mengacu ke 1 sumber
	// bisa menggunakan asterisk atau * agar refrence ke tempat lain juga
	// yaitu person1 dan person2

	// 1. boleh gini
	var person3 *dataSekian = &person1
	person3 = &dataSekian{"gagaga","00","aoisnda"}

	// 2. boleh gini juga
	// person3 := &person1
	// *person3 = dataSekian{"gagaga","00","aoisnda"}

	fmt.Println(person1)
	fmt.Println(person2)
	fmt.Println(person3)

	// ini sama aja seperti ini
	// 						|------------|			|------------|
	// 						| dataSekian |			| dataSekian |
	// 						|------------|			|------------|
	//	 						|						|
	//	 						|						|
	// person2 ----( & )---- person1 -----( & )----- *person3
	// 0x001   		     	0x001
	// 0x009   		     	0x009						0x009
	
	// ------------------------------------------------------------------------------------
	//kemudian untuk di golang yang baru adaalh menggunakan new untuk pointer
	//berikut adalah contohnya

	person4 := new(dataSekian)
	// atau var person4 *dataSekian = new(dataSekian)
	person5 := person4

	// person5.name = "kukururukuk"
	// atau kalau mau isi langsung bisa menggunakan pointer asterisk*
	*person5 = dataSekian{"kukuruku","49","kaaka"}

	fmt.Println(person4)
	fmt.Println(person5)

	//----------------------------------------------------------------------------------
	// kemudian adalah pointer saat ditaruh di function

	person6 := &dataSekian{"hama","877",""} //beri tanda pointer & untuk ngelink ke func changeDataPerson
	funcChangeDataPerson(person6)

	fmt.Println(person6)// tidak berubah hanya keluar nilai {"hama","877",""} kalau belum dikasih pointer

	//----------------------------------------------------------------------------------
	// kemudian adalah pointer saat ditaruh di mmethod
	person7 := dataSekian{"kala","98","dada"}

	person7.MethoChangeDataPerson()//action modif data address

	fmt.Println(person7) // tidak berubah selama methodnya belum dikasih pointer
}

func funcChangeDataPerson(ds *dataSekian)  {
	ds.address = "lolololoh"
}

func (ds *dataSekian) MethoChangeDataPerson()  {
	ds.address = "kota "+ ds.address
}