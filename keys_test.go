package keyboard

import (
	"log"
	"testing"
)

func Test_exportKeyCode(t *testing.T) {

	for _, eks := range Scan {
		log.Printf("Key: %s, Virtual: %v, Scan:%v\n", eks.name, eks.virtual, eks.scan)
	}
	t.Fail()
}
