package main

import "fmt"

func main(){
	myMap:=map[string]string{
		"ruby":"itu cantik",
		"go":"itu cepat",
		"javascript":"itu fancy",
	}
	value,isAvalaible:=myMap["ruby"]  //mencari value dan boolean
	fmt.Println(value)
	fmt.Println(isAvalaible)
	
}