package notifier_test

import (
	"fmt"
	"testing"

	"github.com/rubengglez/birthday_greetings/internal/notifier"
	"github.com/rubengglez/birthday_greetings/internal/shared"
)

type SpySender struct {
	Email   string
	Subject string
	Body    string
	Calls   int
}

func (s *SpySender) Send(email, subject, body string) {
	s.Email = email
	s.Subject = subject
	s.Body = body
	s.Calls++
}

func (s *SpySender) CalledOnceWith(email, subject, body string) bool {
	return s.Email == email && s.Subject == subject && s.Body == body && s.Calls == 1
}

func TestHappyBirthday(t *testing.T) {
	checkMessage := func(t testing.TB, sender *SpySender, email, subject, body string) {
		t.Helper()
		if sender.CalledOnceWith(email, subject, body) != true {
			t.Errorf("send method was called with %s, %s, %s", email, subject, body)
		}
	}

	t.Run("A happy message is sent to Paco", func(t *testing.T) {
		friend := shared.Friend{
			Email:     "paco@example.com",
			FirstName: "Paco",
		}
		sender := &SpySender{}
		notifier := &notifier.DefaultNotifier{}
		notifier.HappyBirthdayWithSender(sender, friend)
		expectedSubject := "Happy Birthday!"
		expectedBody := fmt.Sprintf("Happy birthday, dear %s!", friend.FirstName)

		checkMessage(t, sender, friend.Email, expectedSubject, expectedBody)
	})
}
