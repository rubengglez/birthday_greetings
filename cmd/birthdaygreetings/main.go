package main

import (
	"os"

	"github.com/rubengglez/birthday_greetings/internal/birthday_greetings"
	"github.com/rubengglez/birthday_greetings/internal/dater"
	"github.com/rubengglez/birthday_greetings/internal/notifier"
	"github.com/rubengglez/birthday_greetings/internal/retriever"
)

func main() {
	fs := os.DirFS("./assets")
	// _, err := os.Stat("../assets")
	// if err != nil {
	// 	if os.IsNotExist(err) {
	// 		// File or directory does not exist
	// 		log.Println("no existe!!!!!")
	// 	} else {
	// 		// Some other error. The file may or may not exist
	// 		log.Println("NOooooo existe!!!!!")
	// 	}
	// }
	retriever := retriever.NewFromFileSystem(fs)
	notifier := &notifier.DefaultNotifier{}
	date := dater.NewDefaultDater()
	birthday_greetings.BirthdayGreetings(retriever, notifier, date)
}
