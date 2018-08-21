package main

import "fmt"

type Ntuser struct {
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

func f(Ntuser *Ntuser) {
	*Ntuser.Name = "Alice"
}

func Pointer() {
	name := "Bob"
	ntuser := &Ntuser{
		ID:   4,
		Name: &name,
	}
	fmt.Printf("Name %s\n", *ntuser.Name)
	f(ntuser)
	fmt.Printf("Name %s\n", *ntuser.Name)

	f := Football{}
	f.Bounce()

	fmt.Printf("nado kosong\n")
	ntuser = &Ntuser{}
	// fmt.Printf("name %s\n", *Ntuser.Name)
	name = "luqman"
	ntuser.Name = &name
	fmt.Printf("name %s\n", *ntuser.Name)
}
