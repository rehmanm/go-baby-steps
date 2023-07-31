package greetings

import (
	"math/rand"
)

func randomFormat() string {

	formats := []string{
		"Hi %v, Welcome!!!",
		"Great to see you, %v!",
		"Hail %v! Well met!",
	}

	return formats[rand.Intn(len(formats))]
}

func testFunction() string {
	return "I am test"
}
