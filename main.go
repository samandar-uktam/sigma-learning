package main

// Factory init example

/*
	func main() {
		sms, err := notifier.NewNotifier("sms")
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		if err := sms.Send("Hello!"); err != nil {
			fmt.Printf("Error: %s", err.Error())
		}

		email, err := notifier.NewNotifier("email")
		if err != nil {
			fmt.Printf("Error: %s", err.Error())
		}
		if err := email.Send("Salom!"); err != nil {
			fmt.Printf("Error: %s", err.Error())
		}

	}
*/

import (
	"fmt"
	"sync"
)

type Logger struct{}

func (l *Logger) Log(msg string) {
	fmt.Println("Log: ", msg)
}

var (
	instance *Logger
	once     sync.Once
)

func GetLogger() *Logger {
	once.Do(func() {
		fmt.Println("Creating a new logger...")
		instance = &Logger{}
	})
	return instance
}

func main() {
	log1 := GetLogger()
	log2 := GetLogger()
	fmt.Println("log1->>", &log1)
	fmt.Println("log2->>", &log2)

	log1.Log("Started application.")
	log1.Log("User signed in.")

	if log1 == log2 {
		fmt.Println("Same logger instance is used.✅")
	} else {
		fmt.Println("Different logger instances.❌")
	}
}
