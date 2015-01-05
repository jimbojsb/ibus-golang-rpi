package main

import (
	"os"
	"propellerhead"
	"github.com/johnlauer/serial"
)

func main() {
	ttyPath := os.Args[1]
	config := &serial.Config{Name: ttyPath, Baud: 9600, RtsOn: true}
	serialPort, _ := serial.OpenPort(config)
	f, _ := os.Create(propellerhead.GetWorkingDir() + "/seriallog")

	for {
		byte := make([]byte, 1)
		serialPort.Read(byte)
		f.Write(byte)
	}
}
