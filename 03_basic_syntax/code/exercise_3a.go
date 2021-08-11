package main

import "fmt"

func main() {
	mySentance := "this is a sentance to iterate over"

	for index, value := range mySentance {
		if index % 2 == 0 {
			fmt.Println(string(value))
		}
	}

}