package keyboard

import (
	"bufio"
	"os"
	"os/exec"
)

// Keyboard represents a virtual keyboard.
type Keyboard struct {
	In chan byte
}

// NewKeyboard constructs a new keyboard.
func NewKeyboard(debug bool) Keyboard {
	return Keyboard{make(chan byte, 1)}
}

// Start to listen to the input.
func (kb *Keyboard) Start() {
	initFlags()

	reader := bufio.NewReader(os.Stdin)
	for {
		b, e := reader.ReadByte()
		if e != nil {
			panic(e)
		}

		kb.In <- b
	}
}

// Shutdown must be called or tty will be left in
// the status we used for this keyboard.
func (kb *Keyboard) Shutdown() {
	resetTTY()
}

//////////////////////

func initFlags() {
	stty("-icanon", "-echo")
}

func resetTTY() {
	stty("sane")
}

func stty(flags ...string) {
	cmd := exec.Command("stty", flags...)
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	err := cmd.Run()
	if err != nil {
		panic(err)
	}
}
