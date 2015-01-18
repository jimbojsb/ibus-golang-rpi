package main

import (
	"sync"
	"propellerhead"
	"os"
)

func main() {

	wait := &sync.WaitGroup{}

	if (len(os.Args) > 1) {
		ttyPath := os.Args[1];
		if _, err := os.Stat(ttyPath); err == nil {
			wait.Add(1)
			propellerhead.IbusDevices().SerialInterface.Listen(ttyPath)
		}
	}
	
	wait.Add(1)
	propellerhead.NewAudioController()

	wait.Add(1)
	propellerhead.ServeApi()

	wait.Wait()
}
