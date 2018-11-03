package main

import (
	"log"

	"github.com/kjbreil/keyboard"
)

func main() {
	bs := 100
	ks := 100
	serverAddress := "localhost:10000"

	kb, _ := keyboard.StringToBurst("2282712111", &ks, &bs)
	err := kb.Server(&serverAddress)
	if err != nil {
		log.Fatalln(err)
	}
}
