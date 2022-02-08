package main

import (
	"fmt"
	"testing"
	"time"
)

func hello() {
	fmt.Println("assdfdsf")
}

type TestRetriever struct{}
type TestNotifier struct{}

var x string = ""

func (n TestNotifier) Notify(message string) {
	x = message
}

func (r TestRetriever) Friends() Amigos {
	amigos := make([]Friend, 0)

	friend := Friend{
		LastName:    "Merlo",
		FirstName:   "Manolo",
		DateOfBirth: time.Now(),
		email:       "example@dot.com",
	}
	amigos = append(amigos, friend)
	return amigos
}

func TestBirthdayGreetings(t *testing.T) {
	retriever := TestRetriever{}
	notifier := TestNotifier{}
	expected := "ola k ase"

	BirthdayGreetings(retriever, notifier)

	if x != expected {
		t.Errorf("error got: %s and expected: %s", x, expected)
	}
}
