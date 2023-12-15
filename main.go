package main

import (
	"The-Lebron-James-Notifier/notifier"
	"os"
)

func main() {
	if len(os.Args) < 3 {
		os.Exit(2) // Explained in README.md
	}

	skin := os.Args[1]
	apiKey := os.Args[2]

	notifier := notifier.NewNotifier(skin, apiKey)
	notifier.Run()
}
