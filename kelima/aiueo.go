package main

import (
	"fmt"
)

func main(){
	

	title:="golang itu mantap"

	
	for index,letter:=range title{
		letterString:=string(letter)
		
		if letterString=="a"||letterString=="i"||letterString=="u"||letterString=="e"||letterString=="o"{
			fmt.Println("index:",index,"letter",letterString)
		}
		
	}

}