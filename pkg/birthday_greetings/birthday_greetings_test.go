package birthday_greetings_test

import (
	"testing"

	"github.com/rubengglez/birthday_greetings/pkg/birthday_greetings"
)

type TestRetriever struct{}
type TestNotifier struct {
	message string
	who     string
	Calls   int
}

func (n *TestNotifier) Notify(message, who string) {
	n.message = message
	n.who = who
	n.Calls++
}

func (r *TestRetriever) Friends() birthday_greetings.Amigos {
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
	checkMessage := func(t testing.TB, message, expectedMessage string) {
		t.Helper()
		if message != expectedMessage {
			t.Errorf("got %s want %s", message, expectedMessage)
		}
	}
	checkReceiver := func(t testing.TB, who, expected string) {
		t.Helper()
		if who != expected {
			t.Errorf("got %s want %s", who, expected)
		}
	}
	assertNobodyWasNotifiied := func(t testing.TB, n *TestNotifier) {
		t.Helper()
		if n.Calls != 0 {
			t.Errorf("Someone was notified: %q", n.who)
		}
	}

	t.Run("Manolo has to be notified", func(t *testing.T) {
		retriever := &TestRetriever{}
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 10,
			PlainDay:   8,
		}
		birthday_greetings.BirthdayGreetings(retriever, notifier, date)

		checkMessage(t, notifier.message, "ola k ase")
		checkReceiver(t, notifier.who, "example@dot.com")
	})

	t.Run("No one should be notified as if not his birthday", func(t *testing.T) {
		retriever := &TestRetriever{}
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 1200000,
			PlainDay:   1000000,
		}

		birthday_greetings.BirthdayGreetings(retriever, notifier, date)

		assertNobodyWasNotifiied(t, notifier)
	})
}
