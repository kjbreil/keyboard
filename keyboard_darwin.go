package keyboard

import (
	"fmt"
)

func downKey(key rune) error {
	fmt.Printf("DOWN: %v\n", string(key))
	return nil
}

func upKey(key rune) error {
	fmt.Printf("UP: %v\n", string(key))
	return nil
}

// ListWindowNames is a STUB
func ListWindowNames() error {
	return nil
}

// SetForegroundWindow is a STUB
func SetForegroundWindow(hwnd uintptr) bool {
	return false
}

// FindWindow is a stub
func FindWindow(win string) (ret uintptr, err error) {
	return 1, nil
}
