package main

import (
	"fmt"

	"github.com/hanjoes/keyboard"
)

func main() {
	kb := keyboard.NewKeyboard(false)
	defer kb.Shutdown()
	go func() {
		kb.Start()
	}()

	for {
		select {
		case input := <-kb.In:
			fmt.Printf("keyboard input: %v\n", input.Input)
			if len(input.Input) == 1 && input.Input[0] == 3 {
				return
			}
		}
	}
}
