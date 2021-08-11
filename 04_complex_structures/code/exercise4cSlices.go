package main

import "fmt"


var initialGroceries = []string{"bread", "cheese", "milk", "yerbas", "chicken"}

func addGroceryToList(newGroceries ...string) []string  {
	foods := initialGroceries
	for _, g := range newGroceries {
		foods = append(foods, g)
	}
	return foods
}

func main()  {
	
	grocerieList := addGroceryToList("beets", "beer")
	fmt.Println(grocerieList)
}