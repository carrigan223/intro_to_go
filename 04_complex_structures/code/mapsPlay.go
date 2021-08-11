package main

import "fmt"

func main() {
	// //initializing then declaring
	// var userEmails map[int] string = make(map[int]string)

	// userEmails[1] = "user1@gmail.com"
	// userEmails[2] = "user2@gmail.com"

	// fmt.Println(userEmails)

	//declaring at initializations
	userEmails := map[int] string{
		1: "user1@gmail.com",
		2: "user2@gmail.com",
	}

	//changing the value at 1
	userEmails[1] = "newUser1@gmail.com"
	firstEmail := userEmails[1]

	//checking if a user exists at key 2
	if _, ok := userEmails[2]; ok {
		fmt.Println("It Exists!")
	} else {
		fmt.Println("it does not exist")
	}

	//logging the first user
	fmt.Println(firstEmail)

	//deleteing the user at key 2
	delete(userEmails, 2)
	//logging the map
	fmt.Println(userEmails)

	//checking if user 2 still exists
	if _, ok := userEmails[2]; ok {
		fmt.Println("It Exists!")
	} else {
		fmt.Println("it does not exist")
	}

}