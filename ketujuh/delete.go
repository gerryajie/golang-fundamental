package main

import "fmt"

func main(){
	myMap:=map[string]string{
		"ruby":"itu cantik",
		"go":"itu cepat",
		"javascript":"itu fancy",
	}
	for key,value:=range myMap{
		fmt.Println("key :",key,",","value :",value)
	}

	fmt.Println("==============")

	delete(myMap,"ruby")
	for key,value:=range myMap{
		fmt.Println("key :",key,",","value :",value)
	}
	
}