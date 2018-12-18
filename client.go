package main

import (
	"fmt"
	"github.com/tarcisiodarocha/tmqp"
	"os"
)

func main() {
	conn, err := tmqp.Connect(":5834")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	_, err = conn.QueueDeclare("tarcisio", false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("OK", conn)
}
