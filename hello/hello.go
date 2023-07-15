package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Gladys", "Samantha", "Darrin"}

	messages, err := greetings.Hellos(names)

	// Request a greeting message with no name...
	// -->  message, err := greetings.Hello("")
	// --> message, err := greetings.Hello("Gladys")
	// If an error was returned, print it to the console and
	// exit the program.
	if err != nil {
		// Fatal stops program running.
		log.Fatal(err)
	}

	fmt.Println(quote.Go())

	fmt.Println(messages)
}
