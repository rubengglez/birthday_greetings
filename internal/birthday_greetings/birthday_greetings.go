package birthday_greetings

import (
	"log"
	"time"

	"github.com/rubengglez/birthday_greetings/internal/notifier"
	"github.com/rubengglez/birthday_greetings/internal/retriever"
)

type Date interface {
	Month() int
	Day() int
}

type Birthday struct {
	Month int
	Day   int
}

func BirthdayGreetings(retriever retriever.Retriever, n notifier.Notifier, date Date) {
	friends := retriever.Friends()
	for _, friend := range friends {
		birthday := birthday(friend.DateOfBirth)
		if birthday.Day == date.Day() && birthday.Month == date.Month() {
			log.Println("voy a notificar a", friend)
			n.HappyBirthday(friend)
			continue
		}
		log.Println("NO fue notificado", friend)
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
