package main

import (
	"The-Lebron-James-Notifier/notifier"
	"os"
)

func main() {
	skin := os.Args[1]
	apiKey := os.Args[2]

	notifier := notifier.NewNotifier(skin, apiKey)
	notifier.Run()
}
