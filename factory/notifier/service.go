package notifier

import "fmt"

type SMSNotifier struct{}

func (s *SMSNotifier) Send(message string) error {
	fmt.Println("Sending SMS: ", message)
	return nil
}

type EmailNotifier struct{}

func (e *EmailNotifier) Send(message string) error {
	fmt.Println("Sending email: ", message)
	return nil
}
