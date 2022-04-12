package main

import "fmt"

func main(){
	

	title:="golang itu mantap"

	
	for index,letter:=range title{
		letterString:=string(letter)
		switch letterString{
		case "a","i","u","e","o":
			fmt.Println("index:",index,"letter :",letterString)
		}
		
	}

}