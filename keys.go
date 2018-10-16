package keyboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type keyMap map[rune]int

// The key modifier constans
const (
	Shift        = 0x10
	Ctrl         = 0x11
	Alt          = 0x12
	LeftShift    = 0xA0
	RightShift   = 0xA1
	LeftControl  = 0xA2
	RightControl = 0xA3
)

func exportKeyCode() {
	var km keyMap
	km = make(map[rune]int)
	km['0'] = 0x30
	km['1'] = 0x31
	jsonString, _ := json.Marshal(km)

	fmt.Println(km)
	fmt.Println(jsonString)

	err := ioutil.WriteFile("output.json", jsonString, 0644)

	if err != nil {
		log.Panic(err)
	}

}
