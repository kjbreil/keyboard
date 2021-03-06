// Package keyboard emulates key presses in windows using user32.dll
package keyboard

import (
	"fmt"
	"strings"
	"time"
	"unicode"
)

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
	Sleep   *int `json:",omitempty"`
}

// KeyPress is a single key press
type KeyPress struct {
	Key      rune  `json:",omitempty"`
	Modifier *rune `json:",omitempty"`
	Upper    bool  `json:",omitempty"`
	Sleep    *int  `json:",omitempty"`
}

// StringToBurst takes a string and sleep times and creates a keyburst
func StringToBurst(s string, keySleep *int, burstSleep *int) (b KeyBurst, err error) {
	for _, r := range s {
		var k KeyPress
		if unicode.IsUpper(r) {
			k.Upper = true
		} else {
			r = rune(strings.ToUpper(string(r))[0])
		}
		k.Key = r
		if keySleep != nil {
			k.Sleep = keySleep
		}
		b.Presses = append(b.Presses, k)
	}
	if burstSleep != nil {
		b.Sleep = burstSleep
	}

	return
}

func stringToPress(s string) (k KeyPress, err error) {

	if len(s) > 0 {
		err = fmt.Errorf("Strings longer than a character need to be converted to a burst")
		return
	}

	// if its uppercase set upper flag, otherwise dont set flag but set string
	// to upper case for conversion to rune
	if unicode.IsUpper(rune(s[0])) {
		k.Upper = true
	} else {
		s = strings.ToUpper(s)
	}

	key := keyToRune(s)
	k.Key = key

	// upper case check and addition

	return k, nil

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

// Press a burst of key
func (b *KeyBurst) Press() error {
	for _, kp := range b.Presses {
		err := kp.Press()
		if err != nil {
			return err
		}
	}

	if b.Sleep != nil {
		time.Sleep(time.Duration(*b.Sleep) * time.Millisecond)
	}

	return nil
}

// Press a sequence of keys
func (ks *KeySeq) Press() error {
	for _, kb := range ks.Bursts {
		err := kb.Press()
		if err != nil {
			return err
		}
	}
	return nil
}

// Press presses a single key, just the key
// needs and error return and all that stuff
func Press(key rune) {
	var kp KeyPress
	kp.Key = key
	kp.Press()
	return
}
