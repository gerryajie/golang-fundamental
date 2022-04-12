package main

import "fmt"


func main(){
	scores:=[8]int{100,80,78,92,67,89,60,79}
	var goodScores []int

	for _,score:=range scores{
		if score>=90{
			goodScores=append(goodScores,score)
		}
	}
	fmt.Println(goodScores)
}