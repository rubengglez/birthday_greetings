package birthday_greetings

import (
	"time"
)

type Friend struct {
	LastName    string
	FirstName   string
	DateOfBirth string
	Email       string
}

type Amigos []Friend

type Retriever interface {
	Friends() Amigos
}

type Notifier interface {
	Notify(message, who string)
}

type Date interface {
	Month() int
	Day() int
}

type Birthday struct {
	Month int
	Day   int
}

func BirthdayGreetings(retriever Retriever, notifier Notifier, date Date) {
	friends := retriever.Friends()
	for _, friend := range friends {
		birthday := birthday(friend.DateOfBirth)
		if birthday.Day == date.Day() && birthday.Month == date.Month() {
			notifier.Notify("ola k ase", friend.Email)
		}
	}
}

func birthday(dateOfBirth string) Birthday {
	when, err := time.Parse("2006/01/02", dateOfBirth)

	if err != nil {
		panic(err)
	}

	return Birthday{
		Month: int(when.Month()),
		Day:   when.Day(),
	}
}
