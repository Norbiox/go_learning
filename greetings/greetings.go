package greetings

import (
	"errors"
	"fmt"
)

// Greet returns a greeting for the named person.
func Greet(name string) (string, error) {
	if name == "" {
		return "", errors.New("empty name")
	}
	return fmt.Sprintf("Hello, %v", name), nil
}
