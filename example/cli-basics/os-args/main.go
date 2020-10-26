package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	if len(os.Args) < 2 {
		fmt.Println("no command provided")
		os.Exit(1)
	}

	cmd := os.Args[1]

	switch cmd {
	case "greet":
		greetCmd := flag.NewFlagSet("greet", flag.ExitOnError)

		msgFlag := greetCmd.String("msg", "CLI BASICS", "print hte given value")

		err := greetCmd.Parse(os.Args[2:])
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Printf("hello %s", *msgFlag)

	case "help":

	default:
		fmt.Println("Unknown command given")
	}
}
