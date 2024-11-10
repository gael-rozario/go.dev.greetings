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

    for i:= range name{
        if name[i] == ""{
            return nil, errors.New("Name cannot be empty")
        }
    }

    for i:= range name{
       greetings[name[i]]= fmt.Sprintf(get_greeting_format(),name[i])
    }
    return greetings,nil
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
