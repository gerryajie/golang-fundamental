package main

import "fmt"

func main(){
	var gaming []string
	gaming=append(gaming, "PS4")
	gaming = append(gaming, "Xbox")
	gaming = append(gaming, "SEGA")
	fmt.Println(gaming)

	// gaming:=[]string{"PS4","Xbox","sega"}
	// fmt.Println(gaming)

	for _,console:=range gaming{
		fmt.Println(console)
	}
}