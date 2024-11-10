package greetings

import (
	"fmt"
	"time"
)

func Greet(name string) string {
	message := fmt.Sprintf("Hi there %s! How are you doing\nHappy %s", name, time.Now().Weekday())
	return message
}
