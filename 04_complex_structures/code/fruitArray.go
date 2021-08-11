package main

import "fmt"

func main() {
	fruitArray := [5]string{"apple","bannana","pear","kiwi","peach"}

	splicedFruit := fruitArray[1:3]
	fruitToAdd := append(splicedFruit, "dragon fruit", "mango", "passion fruit", "pineapple", "strawberry", "coconut", "plum")
	fmt.Println(splicedFruit, fruitToAdd)
	fmt.Println(len(fruitToAdd))
	fmt.Println(cap(fruitToAdd))
}