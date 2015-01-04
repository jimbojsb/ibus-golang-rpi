package main

import (
	"propellerhead"
	"os"
	"sync"
)

func main() {
	wait := &sync.WaitGroup{}

	wait.Add(1)
	propellerhead.IbusDevices().SerialInterface.Listen(os.Args[1])
	wait.Wait()
}
