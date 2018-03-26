package main

import (
	"fmt"

	"github.com/hanjoes/keyboard"
)

func main() {
	kb := keyboard.NewKeyboard(false)
	go func() {
		kb.Start()
		defer kb.Shutdown()
	}()

	for {
		select {
		case input := <-kb.In:
			fmt.Printf("keyboard input: %q\n", input)
			break
		}
	}
}
