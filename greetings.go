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

func Greetings(name []string) (map[string]string, error) {

	greetings := make(map[string]string)

	for i := range name {
		greeting, err := Greet(name[i])
		if err != nil {
			return nil,errors.New("name should not be empty")
		}
        greetings[name[i]] = greeting
	}
	return greetings, nil
}

func get_greeting_format() string {

	formatlist := []string{
		"Hi there %s have a great day",
		"How you doing %s",
		fmt.Sprintf("Hi there %%s have a great %s", time.Now().Weekday()),
		"Hello %s welcome to go programming ",
	}
	return formatlist[rand.Intn(len(formatlist))]
}
