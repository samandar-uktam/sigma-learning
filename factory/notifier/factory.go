package notifier

import "fmt"

// NewNotifier is a factory that generates SMSNotifier and EmailNotifier objects.

func NewNotifier(channelType string) (Notifier, error) {
	switch channelType {
	case "sms":
		return &SMSNotifier{}, nil
	case "email":
		return &EmailNotifier{}, nil
	default:
		return nil, fmt.Errorf("unknown notifier type: %s", channelType)
	}
}
