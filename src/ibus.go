package main

import (
	"propellerhead/audio"
	"propellerhead/ibus"
)

func main() {
	ac := new(audio.Controller)
	ibusInterface := ibus.NewInterface(ac)
	ibusInterface.Listen("/dev/ttyUSB0")
}
