package keyboard

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
)

var registration = make(map[byte]Handler)

// Handler is an alias for a handler function that
// can be registered for a specific key.
type Handler = func(*Keyboard, byte)

// Keyboard represents a virtual keyboard.
type Keyboard struct {
	In      chan Input
	kbuffer []byte
	debug   bool
}

// Input represents a keyboard input.
type Input struct {
	Content []byte
}

// NewKeyboard constructs a new keyboard.
func NewKeyboard(debug bool) Keyboard {
	return Keyboard{make(chan Input, 1), make([]byte, 0, 1024), debug}
}

// Start to listen to the input.
func (kb *Keyboard) Start() {
	kb.initKeys()
	initFlags()

	reader := bufio.NewReader(os.Stdin)
	for {
		b, e := reader.ReadByte()
		if e != nil {
			panic(e)
		}
		kb.kbuffer = append(kb.kbuffer, b)

		f := registration[b]
		if f != nil {
			f(kb, b)
		} else if kb.debug {
			echo(kb, b)
		}
		// fmt.Print(b)
	}
}

// Shutdown must be called or tty will be left in
// the status we used for this keyboard.
func (kb *Keyboard) Shutdown() {
	resetTTY()
}

// Register a function to be called when we receive
// the specified byte from keyboard.
func (kb *Keyboard) Register(b byte, f Handler) {
	registration[b] = f
}

func (kb *Keyboard) initKeys() {
	registration[127] = backspace
	registration[10] = remember
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

//////////////////////

func echo(kb *Keyboard, b byte) {
	fmt.Print(string(b))
}

func backspace(kb *Keyboard, b byte) {
	fmt.Print("\b  \b\b")
	kb.kbuffer = kb.kbuffer[0 : len(kb.kbuffer)-2]
}

func remember(kb *Keyboard, b byte) {
	kb.kbuffer = kb.kbuffer[0 : len(kb.kbuffer)-1]
	fmt.Print(string(b))
	// fmt.Print("remembering:\n")
	// fmt.Print(string(kb.kbuffer))
	kb.In <- Input{kb.kbuffer}
	kb.kbuffer = kb.kbuffer[:0]
}
