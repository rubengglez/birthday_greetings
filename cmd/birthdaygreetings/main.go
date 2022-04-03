package main

import (
	"github.com/rubengglez/birthday_greetings/internal/birthday_greetings"
	"github.com/rubengglez/birthday_greetings/internal/dater"
	"github.com/rubengglez/birthday_greetings/internal/notifier"
	"github.com/rubengglez/birthday_greetings/internal/retriever"
)

func main() {
	retriever := &retriever.MemoryRetriever{}
	notifier := &notifier.DefaultNotifier{}
	date := dater.NewDefaultDater()
	birthday_greetings.BirthdayGreetings(retriever, notifier, date)
}
