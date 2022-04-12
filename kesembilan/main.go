package main

import "fmt"

func main(){
	// hitung rata rata
	scores:=[8]int{100,80,78,92,67,89,60,79}

	var total int

	for _,score:=range scores{
		total=total+score
	}

	fmt.Println(total)

	lenght:=len(scores)
	average:=float64(total)/float64(lenght)

	fmt.Println(average)

}