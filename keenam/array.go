package main

import (
	"fmt"
)

func main(){
	languages :=[...]string{"go","java","nodejs","c#","python","kotlin"}
	fmt.Println(languages)
	length:=len(languages)
	fmt.Println(length)
}