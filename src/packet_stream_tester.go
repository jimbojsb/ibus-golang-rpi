package main

import (
	"strings"
	"io/ioutil"
	"encoding/hex"
	"os"
	"propellerhead"
	"github.com/johnlauer/serial"
)

func main() {
	ttyPath := os.Args[1]
	config := &serial.Config{Name: ttyPath, Baud: 9600, RtsOn: true}
	serialPort, _ := serial.OpenPort(config)

	fileContents, _ := ioutil.ReadFile(propellerhead.GetWorkingDir() + "/testcase")
	hexBytes := strings.Split(string(fileContents), " ")
	for _, el := range hexBytes {
		byte, _ := hex.DecodeString(el)
		serialPort.Write(byte)
	}
}
