package main

import "fmt"

func main (){
	cards := []int{0,1,2,3,4,5,6,7,8,9,10}
	
	for _, card := range cards{
		if card % 2 == 0{
			fmt.Println(cards, "%v is even!")
		}else{
			fmt.Println(cards, "%v is odd!")
		}
	}
	
}

