package keyboard

import (
	"log"
	"testing"
)

func Test_exportKeyCode(t *testing.T) {

	for _, eks := range Scan {
		log.Printf("Key: %s, Virtual: %v, Scan:%v\n", eks.Name, eks.Virtual, eks.Scan)
	}
	t.Fail()
}
