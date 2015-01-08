package main

import (
	"os"
	"propellerhead"
	"github.com/mikepb/serial"
	"time"
)

func main() {
	ttyPath := os.Args[1]

	config := serial.RawOptions
	config.FlowControl = serial.FLOWCONTROL_RTSCTS
	config.BitRate = 9600

	serialPort, _ := config.Open(ttyPath)
	serialPort.SetReadDeadline(time.Time{})


	f, _ := os.Create(propellerhead.GetWorkingDir() + "/seriallog")

	for {
		byte := make([]byte, 1)
		serialPort.Read(byte)
		f.Write(byte)
	}
}
