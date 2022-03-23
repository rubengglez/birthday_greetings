package main

import (
	"github.com/rubengglez/birthday_greetings/internal/birthday_greetings"
	"github.com/rubengglez/birthday_greetings/internal/notifier"
	"github.com/rubengglez/birthday_greetings/internal/retriever"
)

type Date struct{}

func (d *Date) Month() int {
	return 10
}

func (d *Date) Day() int {
	return 8
}

func main() {
	retriever := &retriever.MemoryRetriever{}
	notifier := &notifier.DefaultNotifier{}
	date := &Date{}
	birthday_greetings.BirthdayGreetings(retriever, notifier, date)
}
