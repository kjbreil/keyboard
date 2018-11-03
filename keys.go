package keyboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type keyMap map[rune]int

// VirtScan is the Virtual and Scan codes for a key
type VirtScan struct {
	Name    string
	Virtual uint16
	Scan    uint16
}

// Scan is key Scan code map for rune to Scan code
var Scan = map[rune]VirtScan{
	'0': VirtScan{
		Name:    "0",
		Virtual: 0x30,
		Scan:    0x0b,
	},
	'1': VirtScan{
		Name:    "1",
		Virtual: 0x31,
		Scan:    0x02,
	},
	'2': VirtScan{
		Name:    "2",
		Virtual: 0x32,
		Scan:    0x03,
	},
	'3': VirtScan{
		Name:    "3",
		Virtual: 0x33,
		Scan:    0x04,
	},
	'4': VirtScan{
		Name:    "4",
		Virtual: 0x34,
		Scan:    0x05,
	},
	'5': VirtScan{
		Name:    "5",
		Virtual: 0x35,
		Scan:    0x06,
	},
	'6': VirtScan{
		Name:    "6",
		Virtual: 0x36,
		Scan:    0x07,
	},
	'7': VirtScan{
		Name:    "7",
		Virtual: 0x37,
		Scan:    0x08,
	},
	'8': VirtScan{
		Name:    "8",
		Virtual: 0x38,
		Scan:    0x09,
	},
	'9': VirtScan{
		Name:    "9",
		Virtual: 0x39,
		Scan:    0x0a,
	},
	'B': VirtScan{
		Name:    "B",
		Virtual: 0x42,
		Scan:    0x30,
	},
	0x70: VirtScan{
		Name:    "F1",
		Virtual: 0x70,
		Scan:    0x3b,
	},
	0x71: VirtScan{
		Name:    "F2",
		Virtual: 0x71,
		Scan:    0x3c,
	},
	0x72: VirtScan{
		Name:    "F3",
		Virtual: 0x72,
		Scan:    0x3d,
	},
	0x73: VirtScan{
		Name:    "F4",
		Virtual: 0x73,
		Scan:    0x3e,
	},
	0x74: VirtScan{
		Name:    "F5",
		Virtual: 0x74,
		Scan:    0x3f,
	},
	0x75: VirtScan{
		Name:    "F6",
		Virtual: 0x75,
		Scan:    0x40,
	},
	0x76: VirtScan{
		Name:    "F7",
		Virtual: 0x76,
		Scan:    0x41,
	},
	0x77: VirtScan{
		Name:    "F8",
		Virtual: 0x77,
		Scan:    0x42,
	},
	0x78: VirtScan{
		Name:    "F9",
		Virtual: 0x78,
		Scan:    0x43,
	},
	0x79: VirtScan{
		Name:    "F10",
		Virtual: 0x79,
		Scan:    0x44,
	},
	0x7A: VirtScan{
		Name:    "F11",
		Virtual: 0x7A,
		Scan:    0x0a,
	},
	0xA0: VirtScan{
		Name:    "Left Shift",
		Virtual: 0xA0,
		Scan:    0x2a,
	},
	0xA2: VirtScan{
		Name:    "Left Control",
		Virtual: 0xA2,
		Scan:    0x1d,
	},
	0x0D: VirtScan{
		Name:    "Enter",
		Virtual: 0x0D,
		Scan:    0x1c,
	},
	'.': VirtScan{
		Name:    ".",
		Virtual: 0x6E,
		Scan:    0x34,
	},
}

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
