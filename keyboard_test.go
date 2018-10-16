// Package keyboard emulates key presses in windows using user32.dll

package keyboard

import (
	"testing"
)

func TestKeyPress_Press(t *testing.T) {
	// p, _ := stringToPress("a")

	kb, _ := stringToBurst("2282712111")

	kb.Press()

	t.Fail()

	// Press Delete to clear the a from whatever just took the input

	// var s = 6000
	// p = KeyPress{
	// 	Key:   0x08,
	// 	Sleep: &s,
	// }

	// p.Press()
}

func Test_stringToBurst(t *testing.T) {
	// NOTHING
}
