package keyboard

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

type keyMap map[rune]int

// VirtScan is the virtual and scan codes for a key
type VirtScan struct {
	name    string
	virtual uint16
	scan    uint16
}

// Scan is key scan code map for rune to scan code
var Scan = map[rune]VirtScan{
	'0': VirtScan{
		name:    "0",
		virtual: 0x30,
		scan:    0x0b,
	},
	'1': VirtScan{
		name:    "1",
		virtual: 0x31,
		scan:    0x02,
	},
	'2': VirtScan{
		name:    "2",
		virtual: 0x32,
		scan:    0x03,
	},
	'3': VirtScan{
		name:    "3",
		virtual: 0x33,
		scan:    0x04,
	},
	'4': VirtScan{
		name:    "4",
		virtual: 0x34,
		scan:    0x05,
	},
	'5': VirtScan{
		name:    "5",
		virtual: 0x35,
		scan:    0x06,
	},
	'6': VirtScan{
		name:    "6",
		virtual: 0x36,
		scan:    0x07,
	},
	'7': VirtScan{
		name:    "7",
		virtual: 0x37,
		scan:    0x08,
	},
	'8': VirtScan{
		name:    "8",
		virtual: 0x38,
		scan:    0x09,
	},
	'9': VirtScan{
		name:    "9",
		virtual: 0x39,
		scan:    0x0a,
	},
	'B': VirtScan{
		name:    "B",
		virtual: 0x42,
		scan:    0x30,
	},
	0x70: VirtScan{
		name:    "F1",
		virtual: 0x70,
		scan:    0x3b,
	},
	0x71: VirtScan{
		name:    "F2",
		virtual: 0x71,
		scan:    0x3c,
	},
	0x72: VirtScan{
		name:    "F3",
		virtual: 0x72,
		scan:    0x3d,
	},
	0x73: VirtScan{
		name:    "F4",
		virtual: 0x73,
		scan:    0x3e,
	},
	0x74: VirtScan{
		name:    "F5",
		virtual: 0x74,
		scan:    0x3f,
	},
	0x75: VirtScan{
		name:    "F6",
		virtual: 0x75,
		scan:    0x40,
	},
	0x76: VirtScan{
		name:    "F7",
		virtual: 0x76,
		scan:    0x41,
	},
	0x77: VirtScan{
		name:    "F8",
		virtual: 0x77,
		scan:    0x42,
	},
	0x78: VirtScan{
		name:    "F9",
		virtual: 0x78,
		scan:    0x43,
	},
	0x79: VirtScan{
		name:    "F10",
		virtual: 0x79,
		scan:    0x44,
	},
	0x7A: VirtScan{
		name:    "F11",
		virtual: 0x7A,
		scan:    0x0a,
	},
	0xA0: VirtScan{
		name:    "Left Shift",
		virtual: 0xA0,
		scan:    0x2a,
	},
	0xA2: VirtScan{
		name:    "Left Control",
		virtual: 0xA2,
		scan:    0x1d,
	},
	0x0D: VirtScan{
		name:    "Enter",
		virtual: 0x0D,
		scan:    0x1c,
	},
	'.': VirtScan{
		name:    ".",
		virtual: 0x6E,
		scan:    0x34,
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
