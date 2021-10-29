package golibgreet

import (
	"fmt"
)

func Hello(name string) string {
	// Returns a greeting embedding the name
	message := fmt.Sprintf("Welcome %v", name)
	return message
}
