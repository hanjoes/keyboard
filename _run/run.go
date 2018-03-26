package main

import (
	"fmt"

	"github.com/hanjoes/keyboard"
)

func main() {
	kb := keyboard.NewKeyboard(true)
	go func() {
		kb.Start()
		defer kb.Shutdown()
	}()

	for {
		select {
		case input := <-kb.In:
			fmt.Printf("got input: %q\n", input.Content)
			break
		}
	}
}
