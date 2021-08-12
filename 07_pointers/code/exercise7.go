package main

import "fmt"

type User struct {
	ID                         int
	FirstName, LastName, Email string
}

func changeEmail(email *string) string {
	*email = "newEmail@gmail.com"
	return *email
}

func main() {
	user1 := User{
		ID:        1,
		FirstName: "Andrew",
		LastName: "Carrigan",
		Email:    "carrigan223@gmail.com",
	}
	locationInMemoryOfUserEmail := &user1.Email

	fmt.Println("Pointers!")
	fmt.Printf("Andrews email is %s\n", user1.Email)
	fmt.Println(locationInMemoryOfUserEmail)
	fmt.Println(changeEmail(locationInMemoryOfUserEmail))
	fmt.Println(user1)
}
