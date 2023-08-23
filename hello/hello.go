package main

import (
	"fmt"
	"log"

	"example.com/greetings"
	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello, Go World! :)")
	fmt.Println(quote.Go())

	log.SetPrefix("greetings: ")
	// shutdown the log time
	// log.SetFlags(0)

	// Get a greeting message from another package and print it.
	message, err := greetings.Hello("Darling")

	if err != nil {
		log.Fatal(err)
	}

	// If no error was returned, print the returned message to the console.
	fmt.Println(message)

	// Get a slice of names, request greeting messages for the names.
	names := []string{"Yang", "Liu", "Hua"}
	messages, errs := greetings.Hellos(names)

	if errs != nil {
		log.Fatal(errs)
	}

	fmt.Println(messages)

}
