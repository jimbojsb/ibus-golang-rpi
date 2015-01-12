package main

import (
	"sync"
	"propellerhead"
	"os"
)

func main() {

	wait := &sync.WaitGroup{}

	wait.Add(1)
	ttyPath := os.Args[1];
	propellerhead.IbusDevices().SerialInterface.Listen(ttyPath)

	wait.Add(1)
	propellerhead.NewAudioController()

	wait.Wait()
}
