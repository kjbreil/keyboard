// Package keyboard emulates key presses in windows using user32.dll

package keyboard

import "testing"

func TestKeyPress_Press(t *testing.T) {
	p := stringToPress("g")

	err := p.Press()
	if err != nil {
		t.Error(err)
	}

	// Press Delete to clear the a from whatever just took the input

	// var s = 6000
	// p = KeyPress{
	// 	Key:   0x08,
	// 	Sleep: &s,
	// }

	// p.Press()
}
