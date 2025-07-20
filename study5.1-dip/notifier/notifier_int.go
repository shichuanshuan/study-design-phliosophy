package notifier

import "fmt"

type Notifier interface {
	Send(message string) error
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) error {
	fmt.Println("Sending Email:", message)
	return nil
}

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS:", message)
	return nil
}
