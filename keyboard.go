// Package keyboard emulates key presses in windows using user32.dll
package keyboard

import (
	"fmt"
	"strings"
	"time"
	"unicode"

	"golang.org/x/sys/windows"
)

// lazy loading the dll for now
var keyEventDLL = windows.NewLazyDLL("user32.dll").NewProc("keybd_event")

// three types
// sequence = a complete grouping of key presses
// burst = a small grouping of key presses
// press = a single key press, this includes the up and down and modifier
// sleep is time in milliseconds to wait after press or burst

// KeySeq combine bursts into a complete sequence to be run
type KeySeq struct {
	Bursts []KeyBurst
}

// KeyBurst is an array of keys that will be run in a burst
type KeyBurst struct {
	Presses []KeyPress
	Sleep   *int
}

// KeyPress is a single key press
type KeyPress struct {
	Key      rune
	Modifier *rune
	Upper    bool
	Sleep    *int
}

// downKey is the down action on a key, this is how modifiers are added to a key
// press
func downKey(key rune) error {
	vkey := key + 0x80
	_, _, err := keyEventDLL.Call(uintptr(key), uintptr(vkey), 0, 0)
	// error return needs to be corrected before here because its always filled
	// even if there is no error
	return err
}

// upKey the opposite of downkey
func upKey(key rune) error {
	vkey := key + 0x80
	// defining keyUp to be more "verbose"
	var keyUp uintptr = 0x0002
	_, _, err := keyEventDLL.Call(uintptr(key), uintptr(vkey), keyUp, 0)
	return err
}

func stringToPress(s string) (k KeyPress) {

	if len(s) > 0 {
		// not a single character so going to do stuff, right now nothing
	}

	// if its uppercase define set upper flag, otherwise dont set flag but
	// set string to upper case for conversion to rune
	if unicode.IsUpper(rune(s[0])) {
		k.Upper = true
	} else {
		s = strings.ToUpper(s)
	}

	key := keyToRune(s)
	k.Key = key

	// upper case check and addition

	return k

}

// this functions purpose falls on its head unless you pass a single character
func keyToRune(s string) rune {
	if len(s) > 1 {
		panic("string is not single character")
	}
	return rune(s[0])
}

// Press is a single key press, with modifier and case (shift)
// error return does not work
func (p KeyPress) Press() error {
	if p.Modifier != nil {
		err := downKey(*p.Modifier)
		// error is returning "The operation completed successfully." :-(
		if err != nil {
			// return err
		}
	}

	if p.Upper {
		err := downKey(0x10)
		if err != nil {
			// return err
		}
	}

	err := downKey(p.Key)
	if err != nil {
		// return err
	}

	err = upKey(p.Key)
	if err != nil {
		// return err
	}

	if p.Upper {
		err := upKey(0x10)
		// error is returning "The operation completed successfully." :-(
		if err != nil {
			// return err
		}
	}

	if p.Modifier != nil {
		err := upKey(*p.Modifier)
		if err != nil {
			// return err
		}
	}

	// if sleep is defined then sleep for that amount
	if p.Sleep != nil {
		time.Sleep(time.Duration(*p.Sleep) * time.Millisecond)
	}

	return nil
}

func (b *KeyBurst) stringToBurst(s string) {
	for _, l := range s {
		p := KeyPress{
			Key: l,
		}
		b.Presses = append(b.Presses, p)
	}
	fmt.Println(b)
	return
}
