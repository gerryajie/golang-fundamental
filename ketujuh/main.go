package main

import "fmt"

func main(){
	var myMap map[string]int
	myMap=map[string]int{}

	myMap["go"]=9
	myMap["java"]=8
	myMap["ruby"]=7
	
	fmt.Println(myMap["ruby"])
}