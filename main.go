package main

import (
	"The-Lebron-James-Notifier/notifier"
	"fmt"
	"os"
)

const (
	helpMessage string = "Please use the following syntax: go run main.go <skin> <apiKey>"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println(helpMessage)
		os.Exit(1)
	}

	skin := os.Args[1]
	apiKey := os.Args[2]

	notifier := notifier.NewNotifier(skin, apiKey)
	notifier.Run()
}
