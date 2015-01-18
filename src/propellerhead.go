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
		propellerhead.Logger().Info("attempting to listen for ius on " + ttyPath)
		if _, err := os.Stat(ttyPath); err == nil {
			wait.Add(1)
			propellerhead.IbusDevices().SerialInterface.Listen(ttyPath)
		} else {
			propellerhead.Logger().Warn(ttyPath + " not found, skipping ibus")
		}
	}

	wait.Add(1)
	propellerhead.NewAudioController()

	wait.Add(1)
	propellerhead.ServeApi()

	wait.Wait()
}
