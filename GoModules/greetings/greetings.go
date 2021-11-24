package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello returns a greeting for the named person.
func Hello(name string) (string, error) {
	// if no name given return error message.
	if name == "" {
		return "", errors.New("empty name")
	}
	// create a message using a random format
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// Hellos returns a map that associates each of the named people with a greeting
func Hellos(names []string) (map[string]string, error) {
	// a map to associate names w/ messages
	messages := make(map[string]string)
	// Loop through the names calling Hello func.
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		// map the messages to name
		messages[name] = message
	}
	return messages, nil
}

// init values for function use.
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a set of greeting messages (selected @ random)
func randomFormat() string {
	// a slice of message formats.
	formats := []string{
		"Hi, %v. Welcome",
		"Good to see you, %v",
		"Hail, %v! Well Met!",
	}

	// returns a random message by getting a random index of formats[].
	return formats[rand.Intn(len(formats))]
}
