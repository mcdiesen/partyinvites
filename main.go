package main

import "fmt"

type Rsvp struct {
	Name, Email, Phone string
	WillAttend         bool
}

// responses new slice of Rsvp struct
var responses = make([]*Rsvp, 0, 10)

func main() {
	// ...
	fmt.Println("TODO: add some features!")
}
