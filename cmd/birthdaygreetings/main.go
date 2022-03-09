package main

import (
	"fmt"

	"github.com/rubengglez/birthday_greetings/pkg/birthday_greetings"
)

type RealRetriever struct{}

type RealNotifer struct{}

type Date struct{}

func (d *Date) Month() int {
	return 10
}

func (d *Date) Day() int {
	return 8
}

func (r *RealRetriever) Friends() birthday_greetings.Amigos {
	amigos := make([]birthday_greetings.Friend, 0)

	friend := birthday_greetings.Friend{
		LastName:    "Merlo",
		FirstName:   "Manolo",
		DateOfBirth: "1982/10/08",
		Email:       "example@dot.com",
	}
	amigos = append(amigos, friend)
	return amigos
}

func (n *RealNotifer) Notify(message, who string) {
	fmt.Printf("A message %q was sent to %q", message, who)
}

func main() {
	retriever := &RealRetriever{}
	notifier := &RealNotifer{}
	date := &Date{}
	birthday_greetings.BirthdayGreetings(retriever, notifier, date)
}
