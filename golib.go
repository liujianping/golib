package golib

import (
	"errors"
	"fmt"
)

func Greetings(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}

func NumberError(num int) (int, error) {
	if num >= 10 {
		return 0, errors.New("num > 10")
	}
	return num, nil
}
