// package main

// import (
// 	"fmt"
// 	// "strings"
// )

// func main() {
// 	//intiating string variable
// 	var name string
// 	//initiating a pointer for the memory location of string value
// 	var namePointer *string

// 	//assigning string variable
// 	name = "Beyonce"
// 	//assigning the memory location of the string varibale
// 	namePointer = &name
// 	//accessing that string value at the memory location
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)
// }

// // ******************************************************

// func main() {
// 	var name string = "Beyonce"
// 	var namePointer *string = &name
// 	var nameValue = *namePointer

// 	fmt.Println("Name:", name)
// 	fmt.Println("Name *:", namePointer)
// 	fmt.Println("Name Value:", nameValue)

// }

// ******************************************************

// func changeName(n *string) {
// 	*n = strings.ToUpper(*n)
// }

// func main() {
// 	name := "Elvis"
// 	changeName(&name)
// 	fmt.Println(name)
// }

// ******************************************************

//Coordinates represents lattitude and longitude
// type Coordinates struct {
// 	X, Y float64
// }

// var c = Coordinates{X: 10, Y: 20}

// func main()  {
// 	pointerCoordinatesMemoryAddress := &c
	
// 	pointerCoordinatesMemoryAddress.X = 200

// 	fmt.Println(c)
// 	c.Y = 40
// 	fmt.Println(c)
// }