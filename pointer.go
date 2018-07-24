package main

import "fmt"

type User struct {
	ID   int
	Name *string
}

type Ball struct {
}

type Football struct {
	Ball
}

func (b *Ball) Bounce() {
	fmt.Println("Ball bouncing..")
}

func (b *Football) Bounce() {
	fmt.Println("Football bouncing..")
}

func f(user *User) {
	*user.Name = "Alice"
}

func Pointer() {
	name := "Bob"
	user := &User{
		ID:   4,
		Name: &name,
	}
	fmt.Printf("Name %s\n", *user.Name)
	f(user)
	fmt.Printf("Name %s\n", *user.Name)

	f := Football{}
	f.Bounce()

	fmt.Printf("nado kosong\n")
	user = &User{}
	// fmt.Printf("name %s\n", *user.Name)
	name = "luqman"
	user.Name = &name
	fmt.Printf("name %s\n", *user.Name)
}
