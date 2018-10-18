// Package keyboard emulates key presses in windows using user32.dll

package keyboard

import (
	"fmt"
	"log"
	"testing"
)

func TestKeyPress_Press(t *testing.T) {
	// p, _ := stringToPress("a")
	// var ks, bs int

	// bs = 100
	// ks = 100

	// kb, _ := StringToBurst("2282712111", &ks, &bs)

	// kb.Press()

	// t.Fail()

	// var kp KeyPress

	// kp.Key = 0x0D

	// kp.Press()

	// Press Delete to clear the a from whatever just took the input

	// var s = 6000
	// p = KeyPress{
	// 	Key:   0x08,
	// 	Sleep: &s,
	// }2282712111

	// p.Press()

	title := "Windows Powershell"

	h, err := FindWindow(title)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found '%s' window: handle=0x%x\n", title, h)

	// ListWindowNames()

	SetForegroundWindow(h)

	var ks, bs int

	bs = 100
	ks = 100

	kb, _ := StringToBurst("2282712111", &ks, &bs)

	kb.Press()
	t.Fail()

}

func Test_stringToBurst(t *testing.T) {
	// NOTHING
}
