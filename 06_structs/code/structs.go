package main

import "fmt"

// User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

//Group represents set of Users
type Group struct {
	role string
	users []User
	newestUser User
	spaceAvailable bool
}

func describeUser(u User) string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func describeGroup(g Group) string {

	if len(g.users) > 2 {
		g.spaceAvailable = false
	}

	desc := fmt.Sprintf("This user group has %d users, the newest user is %s %s. Accepting new users: %t", len(g.users), g.newestUser.FirstName, g.newestUser.LastName, g.spaceAvailable)
	return desc
}

// // User is a user type
// type User struct {
// ID					       int
// FirstName, LastName, Email string
// }

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	rb := User{ID: 1, FirstName: "Derek", LastName: "Henry", Email: "d.henry@gmail.com"}
	rb2 := User{ID: 1, FirstName: "Alvin", LastName: "Kamara", Email: "a.kamara@gmail.com"}
	rb3 := User{ID: 1, FirstName: "Dalvin", LastName: "Cook", Email: "d.Cook@gmail.com"}
	rb4 := User{ID: 1, FirstName: "Chritian", LastName: "McCaffrey", Email: "c.McCaffrey@gmail.com"}

	g := Group{role: "Best RBs", users: []User{rb, rb2, rb3, rb4}, newestUser: rb, spaceAvailable: true}

	fmt.Println(u)
	fmt.Println(u.FirstName)
	fmt.Println(u.LastName)
	fmt.Println(describeUser(u))
	fmt.Println(describeGroup(g))
	fmt.Println(g)

}
