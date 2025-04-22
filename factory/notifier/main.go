package notifier

import "fmt"

// Factory init example

func main() {
	sms, err := NewNotifier("sms")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	if err := sms.Send("Hello!"); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

	email, err := NewNotifier("email")
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	if err := email.Send("Salom!"); err != nil {
		fmt.Printf("Error: %s", err.Error())
	}

}
