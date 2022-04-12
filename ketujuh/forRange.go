package main

import "fmt"

func main(){
	myMap:=map[string]string{
		"ruby":"itu cantik",
		"go":"itu cepat",
	}
	for key,value:=range myMap{
		fmt.Println("key :",key,"value :",value)
	}
	
}