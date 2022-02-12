package main

import (
	"fmt"
	"testing"
)

func hello() {
	fmt.Println("assdfdsf")
}

type TestRetriever struct{}
type TestNotifier struct{}

var x string = ""

func (n *TestNotifier) Notify(message, who string) {
	x = message + " - " + who
}

func (r *TestRetriever) Friends() Amigos {
	amigos := make([]Friend, 0)

	friend := Friend{
		LastName:    "Merlo",
		FirstName:   "Manolo",
		DateOfBirth: "1982/10/08",
		Email:       "example@dot.com",
	}
	amigos = append(amigos, friend)
	return amigos
}

type SpyDate struct {
	PlainMonth int
	PlainDay   int
}

func (d *SpyDate) Month() int {
	return d.PlainMonth
}

func (d *SpyDate) Day() int {
	return d.PlainDay
}

func TestBirthdayGreetings(t *testing.T) {
	t.Run("Manolo has to be notified", func(t *testing.T) {
		x = ""
		retriever := &TestRetriever{}
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 10,
			PlainDay:   8,
		}
		expected := "ola k ase - example@dot.com"

		BirthdayGreetings(retriever, notifier, date)

		if x != expected {
			t.Errorf("error got: %s and expected: %s", x, expected)
		}
	})

	t.Run("No one should be notified as if not his birthday", func(t *testing.T) {
		x = ""
		retriever := &TestRetriever{}
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 1200000,
			PlainDay:   1000000,
		}
		expected := ""

		BirthdayGreetings(retriever, notifier, date)

		if x != expected {
			t.Errorf("error got: %s and expected: %s", x, expected)
		}
	})
}
