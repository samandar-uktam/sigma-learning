package oop

import "fmt"

type Notifier interface {
	SendNotification(msg string) error
}

type EmailNotifier struct {
	Email string
}

func (e EmailNotifier) SendNotification(msg string) error {
	fmt.Printf("Sending email to %s: %s\n", e.Email, msg)
	return nil
}

func Alert(n Notifier) {
	_ = n.SendNotification("Your account has been accessed!")
}

// Here the Alert function is abstracter meaning that it only knows about Notifier interface,
// it doesn't care about how the message is notified.
