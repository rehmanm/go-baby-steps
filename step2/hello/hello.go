package main

/*
 	public vs private function
	all functions started with Capital Letters are public e.g. Hello("Rehman")
	all functions started with Small Letters are private to module e.g. randomFormat()

	https://stackoverflow.com/a/22148240
*/
import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Rehman")
	if err == nil {
		fmt.Println(message)
	} else {
		log.Fatal(err)
	}

}
