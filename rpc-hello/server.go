package main

import "fmt"

type Message string

func (m *Message) Echo (args *string, response *string) (err error) {
	err = nil
	*response = fmt.Sprintf("Echo: %s", *args)
	return
}