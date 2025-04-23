package logger

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

	log1.Log("Started application.")
	log1.Log("User signed in.")

	if log1 == log2 {
		fmt.Println("Same logger instance is used.✅")
	} else {
		fmt.Println("Different logger instances.❌")
	}
}
