package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// set logger props
	// the log entry prefix to disable printing
	// the time, src file + line num
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// a slice of names
	names := []string{"bob", "jane", "cam"}

	// request a greeting
	message, err := greetings.Hellos(names)
	// if error print to console + exit
	if err != nil {
		log.Fatal(err)
	}

	// if no error procede
	// get a greeting message and print it.
	fmt.Println(message)
}
