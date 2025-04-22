package notifier

import (
	"fmt"
)

// interfaces
type EmailSender interface {
	Send(to, message string)
}

type SmsSender interface {
	Send(to, message string)
}

type PushSender interface {
	Send(to, message string)
}

// entities (for production env)
type ProdEmailSender struct{}

func (p *ProdEmailSender) Send(to, message string) {
	fmt.Printf("Sending email message: %s to: %s", message, to)
}

type ProdSmsSender struct{}

func (p *ProdSmsSender) Send(to, message string) {
	fmt.Printf("Sending sms message: %s to: %s", message, to)
}

type ProdPushSender struct{}

func (p *ProdPushSender) Send(to, message string) {
	fmt.Printf("Sending push message: %s to: %s", message, to)
}

// entities (for development env)
type DevEmailSender struct{}

func (p *DevEmailSender) Send(to, message string) {
	fmt.Printf("[Test] Sending email message: %s to: %s", message, to)
}

type DevSmsSender struct{}

func (p *DevSmsSender) Send(to, message string) {
	fmt.Printf("[Test] Sending sms message: %s to: %s", message, to)
}

type DevPushSender struct{}

func (p *DevPushSender) Send(to, message string) {
	fmt.Printf("[Test] Sending push message: %s to: %s", message, to)
}
