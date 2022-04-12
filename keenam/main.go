package main

import "fmt"

func main(){
	var languages [5]string
	languages[0]="go"
	languages[1]="ruby"
	languages[2]="nodejs"
	languages[3]="python"
	languages[4]="java"
	fmt.Println(languages)

	length:=len(languages)
	fmt.Println(length)


}