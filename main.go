package main

type Mahasiswa struct{
	Name string
	age int
}

func NewMahasiswa (name string, age int) *Mahasiswa{
	if age < 20{
		return nil
	}

	return &Mahasiswa{
		Name : Name, 
		Age : age,
	}
}

func main() {

	mahasiswa1 := NewMahasiswa("alex", 19)
	if mahasiswa1 == nil{
		fmt.Printf("ini belum bisa masuk jadi mahasiswa")
	}
	else{
		fmt.Printf("Hello", mahasiswa1.Name)
	}

	toba := NewMahasiswa ("tobat", 21)
	if toba == nil{
	fmt.Printf("ini belum bisa masuk jadi mahasiswa")
	}
	else{
		fmt.Printf("Hello", toba.Name)
	}
}
	