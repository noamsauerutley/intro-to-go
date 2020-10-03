package main

import "fmt"

// User is a user type
// type User struct {
// 	ID                         int
// 	FirstName, LastName, Email string
// }

// // User is a user type
type User struct {
	ID                         int
	FirstName, LastName, Email string
}

// Group represents a set of users
type Group struct {
	Role           string
	Users          []User
	NewestUser     User
	SpaceAvailable bool
}

func (u User) describe() string {
	desc := fmt.Sprintf("Name: %s %s, Email: %s", u.FirstName, u.LastName, u.Email)
	return desc
}

func (g *Group) describe() string {
	if len(g.Users) > 2 {
		g.SpaceAvailable = false
	}
	desc := fmt.Sprintf("This user group has %v users. The newest user is %s %s. Accepting new users? %t", len(g.Users), g.NewestUser.FirstName, g.NewestUser.LastName, g.SpaceAvailable)
	return desc
}

func main() {
	u := User{ID: 1, FirstName: "Marilyn", LastName: "Monroe", Email: "marilyn.monroe@gmail.com"}
	u2 := User{ID: 2, FirstName: "Bob", LastName: "Loblaw", Email: "bob.loblaw@gmail.com"}
	u3 := User{ID: 3, FirstName: "Gob", LastName: "Bluth", Email: "gob@gmail.com"}

	g := Group{
		Role:           "Dog People",
		Users:          []User{u, u2, u3},
		NewestUser:     u3,
		SpaceAvailable: true}

	fmt.Println(u.describe())
	fmt.Println(g.describe())
}
