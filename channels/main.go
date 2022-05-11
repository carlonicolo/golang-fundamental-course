package main

import (
	"fmt"

	"github.com/eiannone/keyboard"
)

// rune single character for build a string
// channel accept only a single type
var keyPressChan chan rune

// in the terminal
// go get github.com/eiannone/keyboard

func main() {

	keyPressChan = make(chan rune)

	go listenForKeyPress()

	fmt.Println("Press any key, or q to quit")

	// open a keyboard and ignore errors, after defer to close the keyboard
	_ = keyboard.Open()
	defer func() {
		keyboard.Close()
	}()

	for {
		char, _, _ := keyboard.GetSingleKey()
		if char == 'q' || char == 'Q' {
			break
		}

		keyPressChan <- char
	}

	/*
		go doSomething("Hello, world")

		fmt.Println("This is another message")
		for {
			// do something
		}
	*/

}

func listenForKeyPress() {
	for {
		key := <-keyPressChan
		fmt.Println("You pressed", string(key))
	}
}

/*
func doSomething(s string) {
	until := 0
	for {
		fmt.Println("s is", s)
		until++
		if until >= 5 {
			break
		}
	}
}
*/
