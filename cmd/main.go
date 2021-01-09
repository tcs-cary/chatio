package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
)

func main() {
	fmt.Println("Welcome to chat.io!")
	var user string
	fmt.Print("What's your name? ")
	_, err := fmt.Scanln(&user)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(" ")
	userMF := "[%s]: %s" // userMessageFormat
	fmt.Printf(userMF, user, "Hello, World")
}
