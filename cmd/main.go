package main

import (
	"fmt"
)

type action struct {
	msg   string
	value string
}

var (
	messages []action
	counter  = 0
)

func init() {
	messages = []action{
		{
			msg:   "key",
			value: "",
		}, {
			msg:   "operation",
			value: "",
		}, {
			msg:   "source",
			value: "",
		}, {
			msg:   "destination",
			value: "",
		},
	}
}
func main() {
	fmt.Println("crypto!!!!,")
	for counter <= len(messages) {
		msg := ""
		fmt.Println(messages[counter].msg)
		_, err := fmt.Scan(&msg)
		if err != nil || msg == "" {
			continue
		}
		messages[counter].value = msg
		counter++
	}

}
