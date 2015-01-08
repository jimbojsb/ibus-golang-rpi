package main

import (
	"strings"
	"io/ioutil"
	"encoding/hex"
	"os"
	"propellerhead"
	"github.com/mikepb/serial",
	"time"
)

func main() {
	ttyPath := os.Args[1]

	config := serial.RawOptions
	config.FlowControl = serial.FLOWCONTROL_RTSCTS
	config.BitRate = 9600

	serialPort, _ := config.Open(ttyPath)
	serialPort.SetReadDeadline(time.Time{})

	fileContents, _ := ioutil.ReadFile(propellerhead.GetWorkingDir() + "/testcase")
	hexBytes := strings.Split(string(fileContents), " ")
	for _, el := range hexBytes {
		byte, _ := hex.DecodeString(el)
		serialPort.Write(byte)
	}
}
