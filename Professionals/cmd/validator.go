/* validator.go */

package main

import (
	"strings"
)

type Message struct {
	Id     int
	Name   string
	Age    string
	Employ string
	Errors map[string]string
}

func (msg *Message) Validate() bool {
	msg.Errors = make(map[string]string)

	if strings.TrimSpace(msg.Name) == "" {
		msg.Errors["Name"] = "Please enter a name."
	}

	if strings.TrimSpace(msg.Age) == "" {
		msg.Errors["Age"] = "Please enter age."
	}

	if strings.TrimSpace(msg.Employ) == "" {
		msg.Errors["Employ"] = "Please enter employment."
	}

	return len(msg.Errors) == 0
}


