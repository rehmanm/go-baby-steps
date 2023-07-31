package greetings

import (
	"errors"
	"fmt"
	"math/rand"
)

func Hellos(names []string) (map[string]string, error) {

	if len(names) == 0 {
		return nil, errors.New("Please provide at least one name")
	}

	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
}

func Hello(name string) (string, error) {
	if name == "" {
		return "", errors.New("Name can't be blank")
	}
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

func randomFormat() string {
	// A slice of message formats.
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you, %v!",
		"Hail, %v! Well met!",
	}

	// Return one of the message formats selected at random.
	return formats[rand.Intn(len(formats))]
}
