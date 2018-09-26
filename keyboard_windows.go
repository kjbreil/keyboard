package keyboard

import (
	"golang.org/x/sys/windows"
)

// lazy loading the dll for now
var keyEventDLL = windows.NewLazyDLL("user32.dll").NewProc("keybd_event")

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
