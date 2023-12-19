package module1

func TypeDeclaration()  {
	
	// buat sejenis func simple dengan type yg bisa dimasukkan parameter
	type nama string

	const namaPanggilan nama = "naufall"

	println(namaPanggilan)

	// type nama string
	println(nama("gagagagagag"))

	// func test(data string)string
	println(test("gagagag"))
}


func test(data string)string{
	return data
}