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
