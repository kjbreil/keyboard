package keyboard

import (
	"fmt"
)

func downKey(key rune) error {
	fmt.Printf("%s", string(key))
	return nil
}

func upKey(key rune) error {
	fmt.Printf("%s", string(key))
	return nil
}
