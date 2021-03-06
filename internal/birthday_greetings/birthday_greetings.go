package birthday_greetings

import (
	"time"

	"github.com/rubengglez/birthday_greetings/internal/dater"
	"github.com/rubengglez/birthday_greetings/internal/notifier"
	"github.com/rubengglez/birthday_greetings/internal/retriever"
)

type birthdayData struct {
	Month int
	Day   int
}

func BirthdayGreetings(retriever retriever.Retriever, n notifier.Notifier, date dater.Date) {
	friends := retriever.Friends()
	for _, friend := range friends {
		birthday := birthday(friend.DateOfBirth)
		if birthday.Day == date.Day() && birthday.Month == date.Month() {
			n.HappyBirthday(friend)
		}
	}
}

func birthday(dateOfBirth string) birthdayData {
	when, err := time.Parse("2006/01/02", dateOfBirth)

	if err != nil {
		panic(err)
	}

	return birthdayData{
		Month: int(when.Month()),
		Day:   when.Day(),
	}
}
