package main

import "fmt"


var initialPets map[string]string = map[string]string{
	"nacho":"dog",
	"karma":"dog",
	"wizard":"turtle",
	"alley":"cat",
	"pepsi":"fish",

}

func exist(petName string) bool {
	_, ok := initialPets[petName]
	return ok
}

func main()  {
	pet := "spot"
	petExists := exist(pet)
	fmt.Println(petExists)
}