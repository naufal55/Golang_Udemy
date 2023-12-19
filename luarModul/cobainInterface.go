package luarmodul

import "fmt"

type kelakuan interface {
	berjalan() string
	berlari() string
	menyerang() string
}

type orang struct {
	nama   string
	darah  int8
	damage int
	isHit bool
	isHitted bool
}

//------------- basic 
func (o orang) berjalan() string{
	return  o.nama+"sedang berjalan"
}

func (o orang) berlari() string {
	return  o.nama+"sedang berlari"
}

func (o orang) menyerang()string  {
	return o.nama+"sedang menyerang"
}

//------------------ skill

func dash(k kelakuan) {
	fmt.Println(k.berlari())
	fmt.Println(k.menyerang())
}

func Action() {
	person1 := orang{
		nama:   "A",
		darah:  100,
		damage: 5,
		isHit: false,
		isHitted: false,
	}
	person2 := orang{
		nama:   "B",
		darah:  100,
		damage: 7,
		isHit: false,
		isHitted: false,
	}

	dash(person1)
	dash(person2)
}


// jadi pada pembahasan interface ini adalah
// saat dash(person) isi dari parameter nya adalah k kelakuan
// yg berisi field interface kelakuan itu sendiri => interface kelakuan
// dan action dash melakukan print nilai orang, sesuai method yg
// diimplent oleh struct orang => menyrang(), berjalan() dan berlari()
// dan akhirnya do action

