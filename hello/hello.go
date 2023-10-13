package main

import (
	"fmt"
	"log"

	"github.com/natanaelrusli/greetings"
)

func main() {
	message, err := greetings.Hello("Gladyss")

	if err != nil {
		log.Fatal(err.Error())
	}

	fmt.Println(message)
}
