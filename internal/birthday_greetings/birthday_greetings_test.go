package birthday_greetings_test

import (
	"reflect"
	"testing"

	"github.com/rubengglez/birthday_greetings/internal/birthday_greetings"
	"github.com/rubengglez/birthday_greetings/internal/shared"
)

type TestRetriever struct {
	FriendsStored shared.Friends
}
type TestNotifier struct {
	Notified shared.Friends
	Calls    int
}

func (n *TestNotifier) HappyBirthday(friend shared.Friend) {
	n.Calls++
	if n.Notified == nil {
		n.Notified = make(shared.Friends, 0)
	}
	n.Notified = append(n.Notified, friend)
}

func (n *TestNotifier) WasNotified(friend shared.Friend) bool {
	if n.Notified == nil {
		return false
	}
	for _, who := range n.Notified {
		if reflect.DeepEqual(who, friend) {
			return true
		}
	}
	return false
}

func (r *TestRetriever) Friends() shared.Friends {
	return r.FriendsStored
}

func (r *TestRetriever) AddFriend(friend shared.Friend) {
	if r.FriendsStored == nil {
		r.FriendsStored = make(shared.Friends, 0)
	}
	r.FriendsStored = append(r.FriendsStored, friend)
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
	checkReceiver := func(t testing.TB, n *TestNotifier, friend shared.Friend) {
		t.Helper()
		if !n.WasNotified(friend) {
			t.Errorf("%v was NOT notified", friend)
		}
	}
	checkWhoNotReceivesMsg := func(t testing.TB, n *TestNotifier, friend shared.Friend) {
		t.Helper()
		if n.WasNotified(friend) {
			t.Errorf("%v was notified when it shouldnt", friend)
		}
	}
	assertNobodyWasNotifiied := func(t testing.TB, n *TestNotifier) {
		t.Helper()
		if n.Calls != 0 {
			t.Errorf("Someone was notified: %q", n.Notified)
		}
	}

	t.Run("Manolo has to be notified when it is his birthday and it is the only one", func(t *testing.T) {
		friend := shared.Friend{
			LastName:    "Merlo",
			FirstName:   "Manolo",
			DateOfBirth: "1982/10/08",
			Email:       "example@dot.com",
		}
		retriever := &TestRetriever{}
		retriever.AddFriend(friend)
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 10,
			PlainDay:   8,
		}
		birthday_greetings.BirthdayGreetings(retriever, notifier, date)

		checkReceiver(t, notifier, friend)
	})

	t.Run("Manolo has to be notified when it is his birthday", func(t *testing.T) {
		friend := shared.Friend{
			LastName:    "Merlo",
			FirstName:   "Manolo",
			DateOfBirth: "1982/10/08",
			Email:       "example@dot.com",
		}
		pepe := shared.Friend{
			LastName:    "Pardexo",
			FirstName:   "Pepe",
			DateOfBirth: "1982/10/10",
			Email:       "example@dot.com",
		}
		retriever := &TestRetriever{}
		retriever.AddFriend(friend)
		retriever.AddFriend(pepe)
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 10,
			PlainDay:   8,
		}
		birthday_greetings.BirthdayGreetings(retriever, notifier, date)

		checkReceiver(t, notifier, friend)
		checkWhoNotReceivesMsg(t, notifier, pepe)
	})

	t.Run("Manolo and Pepe have to be notified because it's their birthday", func(t *testing.T) {
		friend := shared.Friend{
			LastName:    "Merlo",
			FirstName:   "Manolo",
			DateOfBirth: "1982/10/08",
			Email:       "example@dot.com",
		}
		bob := shared.Friend{
			LastName:    "Bob",
			FirstName:   "Bob",
			DateOfBirth: "1982/09/10",
			Email:       "bob@dot.com",
		}
		pepe := shared.Friend{
			LastName:    "Pardexo",
			FirstName:   "Pepe",
			DateOfBirth: "1983/10/08",
			Email:       "pepe@dot.com",
		}
		anne := shared.Friend{
			LastName:    "Mary",
			FirstName:   "Anne",
			DateOfBirth: "1982/10/09",
			Email:       "anne@dot.com",
		}
		retriever := &TestRetriever{}
		retriever.AddFriend(friend)
		retriever.AddFriend(pepe)
		retriever.AddFriend(bob)
		retriever.AddFriend(anne)
		notifier := &TestNotifier{}
		date := &SpyDate{
			PlainMonth: 10,
			PlainDay:   8,
		}
		birthday_greetings.BirthdayGreetings(retriever, notifier, date)

		checkReceiver(t, notifier, friend)
		checkReceiver(t, notifier, pepe)
		checkWhoNotReceivesMsg(t, notifier, anne)
		checkWhoNotReceivesMsg(t, notifier, bob)
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
