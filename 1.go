package main

import (
	"fmt"
)

type Human struct {
	Name string
	Age int
}

func (h *Human) Talk() {
	fmt.Printf("Helllo! I am %s, and i am %d years old\n", h.Name, h.Age)
}

type Action struct {
	Human
}
	
func main() {
	a := Action{Human{"Max", 3}}
	h := Human{"David", 2}
	h.Talk()
	a.Talk()
}