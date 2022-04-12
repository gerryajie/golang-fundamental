package main

import "fmt"

func main(){
	
	title:="golang itu mantap"

	for _,letter:=range title{
		fmt.Println("letter:",string(letter))
	}
	// for index,letter:=range title{
	// 	if index%2==0{
	// 	fmt.Println("index:",index,"letter:",string(letter))
	// 	}
	// }

}