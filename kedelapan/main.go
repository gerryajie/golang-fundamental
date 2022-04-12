package main

import "fmt"

func main(){
	students:=[]map[string]string{
		{"name":"agung","score":"A"},
		{"name":"busara","score":"B"},
		{"name":"loly","score":"E"},

	}
	for _,student:=range students{
		fmt.Println(student["name"],"---",student["score"])
	}
}