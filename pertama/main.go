package main

import (
	"fmt"
	"pertama/calculation"
)
func main(){
	fmt.Println("halo belajar golang")
	sentence:=calculation.TestAja()

	fmt.Println(sentence)
	result:=calculation.Add(15,30)
	resultMultply:=calculation.Multy(3,19)
	resultBagi:=calculation.Bagi(4,2)
	resultKurang:=calculation.Kurang(30,24)



	fmt.Println(result,"Hasil Tambah")
	fmt.Println(resultMultply,"Hasil Kali")
	fmt.Println(resultBagi,"Hasil bagi")
	fmt.Println(resultKurang,"Hasil Kurang")

}


// 1.Executable (Langsung di eksekusi) 
// 2.library(tidak langsung di eksekusi)