package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// If no name was given, return an error with a message.
	if name == "" {
		return name, errors.New("empty name")
	}
	// Create a message using a random format.
	// Correct --> message := fmt.Sprintf(randomFormat(), name)
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hello returns a map that associates each of the named people
// with a greeting message.
func Hellos(names []string) (map[string]string, error) {

	// A map associating names with messages
	messages := make(map[string]string)
	// Loop through, calling Hello to get a message per name.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// In the map, associate message to name
		messages[name] = message
	}
	return messages, nil
}

// init sets intial values for variables used in the function.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns one of a set of greeting messages.
// The returned message is selected at random.
func randomFormat() string {
	// A slice of message formats
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return a random greeting
	return formats[rand.Intn(len(formats))]

}
