package main

import "fmt"

func main() {
	scores := [5]float64{2.0, 3.5, 10.2, 4.7, 12.9}
	for _, value := range scores {
		fmt.Println(value)
	}
	
}