package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

func Greet(name string) (string, error) {

	if name == "" {
		return "", errors.New("name should not be empty")
	}
	message := fmt.Sprintf(get_greeting_format(), name)

	return message, nil
}

func get_greeting_format() string {

	formatlist := []string{
		"Hi there %s have a great day",
		"How you doing %s",
		fmt.Sprintf("Hi there %%s have a great %s", time.Now().Day()),
		// "Hello %s welcom to go programming ",
	}
	return formatlist[rand.Intn(len(formatlist))]
}
