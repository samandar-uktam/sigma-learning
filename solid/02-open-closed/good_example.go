package solid

// This is a good example because it is open for extension, and closed for modification.

// As long as you want to add a new channel to notify you can create a new struct and implement its own method to send the message to that channel.

/*
	type NotificationSender interface {
		Send(message string)
	}

	type EmailSender struct{}

	func (e EmailSender) Send(message string) {
		fmt.Println("📧 Email:", message)
	}

	type SMSSender struct{}

	func (s SMSSender) Send(message string) {
		fmt.Println("📱 SMS:", message)
	}

	type SlackSender struct{}

	func (s SlackSender) Send(message string) {
		fmt.Println("💬 Slack:", message)
	}

	// after some time i am adding a new channel for Telegram.
	func (s TelegramSender) Send(message string) {
		fmt.Println("💬 Telegram:", message)
	}

	type Notifier struct {
		sender NotificationSender
	}

	func (n Notifier) Notify(message string) {
		n.sender.Send(message)
	}

	func main() {
		email := Notifier{sender: EmailSender{}}
		sms := Notifier{sender: SMSSender{}}
		slack := Notifier{sender: SlackSender{}}

		email.Notify("Hello via email!")
		sms.Notify("Hello via SMS!")
		slack.Notify("Hello via Slack!")
	}
*/
