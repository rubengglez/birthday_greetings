package main

import (
	"fmt"
	"time"
)

type Friend struct {
	LastName    string
	FirstName   string
	DateOfBirth time.Time
	email       string
}

type Amigos []Friend

type Retriever interface {
	Friends() Amigos
}

type Notifier interface {
	Notify(message string)
}

func BirthdayGreetings(retriever Retriever, notifier Notifier) {
	notifier.Notify("ola k ase")
}

func main() {
	fmt.Println("First commit")
}
