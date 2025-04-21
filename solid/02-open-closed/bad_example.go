package solid

// This is a bad example, because Notify method is not extensive. When we want to add a new channel, the method body must be changed.
// At last the Notify method will be messy and difficult to maintain

/*
	type Notifier struct{}

	func (n Notifier) Notify(channel string, message string) {
		if channel == "email" {
			fmt.Println("Sending email:", message)
		} else if channel == "sms" {
			fmt.Println("Sending SMS:", message)
		} else if channel == "slack" {
			fmt.Println("Sending Slack message:", message)
		}
	}
*/
