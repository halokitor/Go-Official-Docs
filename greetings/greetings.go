package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name input, return an error with a message.
	if name == "" {
		return "", errors.New("empty Name was given.:(")
		// return name, errors.New("empty Name was given.:(")
	}

	// If input name was given, return a greeting that embeds the name in a message.
	message := fmt.Sprintf(randomFormat(), name)
	// message := fmt.Sprint(randomFormat())
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting message.
func Hellos(names []string) (map[string]string, error) {
	// A map to associate names with messages.
	messages := make(map[string]string)

	// Loop through the received slice of names
	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}

		// If name was given.
		messages[name] = message
	}
	return messages, nil
}

// randomFormat, returns one of a set of greeting messages.
func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome! :)",
		"Nice to meet u, %v. :>",
		"Hail, %v. Well met!",
	}

	// Must set the rand.Seed, usually use time.Now()
	rand.Seed(time.Now().UnixNano())

	// Get a random num in range len.
	randomNum := rand.Intn(len(formats))

	// Return a randomly selected message format
	return formats[randomNum]
}
