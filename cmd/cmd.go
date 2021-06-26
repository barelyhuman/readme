package cmd

import "fmt"

var message string

func Setup() {
	message = "Hello World"
}

func Run() {
	fmt.Println(message)
}
