package greetings

import (
	"errors"
	"fmt"
	"time"
)

func Greet(name string) (string,error) {

    if name != ""{
        return "",errors.New("name should not be empty")
    }
	message := fmt.Sprintf("Hi there %s! How are you doing\nHappy %s", name, time.Now().Weekday())
	return message,nil
}
