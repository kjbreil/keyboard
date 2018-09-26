package keyboard

import (
	"fmt"
)

func downKey(key rune) error {
	fmt.Printf("DOWN: %s", string(key))
	return nil
}

func upKey(key rune) error {
	fmt.Printf("UP: %s", string(key))
	return nil
}
