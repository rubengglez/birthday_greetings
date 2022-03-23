package notifier

import (
	"fmt"

	"github.com/rubengglez/birthday_greetings/internal/shared"
)

const SUBJECT = "Happy Birthday!"

type Notifier interface {
	HappyBirthday(friend shared.Friend)
}
type Sender interface {
	Send(email, subject, body string)
}

type DefaultNotifier struct{}
type defaultSender struct{}

func (s *defaultSender) Send(email, subject, body string) {}

func (n *DefaultNotifier) HappyBirthday(f shared.Friend) {
	sender := &defaultSender{}
	sender.Send(f.Email, SUBJECT, body(f))
}

func (n *DefaultNotifier) HappyBirthdayWithSender(s Sender, f shared.Friend) {
	s.Send(f.Email, SUBJECT, body(f))
}

func body(f shared.Friend) string {
	return fmt.Sprintf("Happy birthday, dear %s!", f.FirstName)
}
