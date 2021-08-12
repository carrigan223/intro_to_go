package main

import (
	"fmt"
	"fem-intro-to-go/05_toolkit/code/utils"
)

func calculateData() int {
	totalValue := utils.Add(1,2,3,4,5)
	return totalValue
}

func main() {
	fmt.Println("Packages!")
	total := calculateData()
	fmt.Println(total)
	fmt.Println(utils.MakeExcited("my name is andrew"))
}
