package main

import (
	"fmt"
	"os"
	"propellerhead"
	"strings"
	"github.com/johnlauer/serial"
)

func main() {
	ttyPath := os.Args[1];
	config := &serial.Config{Name: ttyPath, Baud: 9600, RtsOn: true}
	serialPort, _ := serial.OpenPort(config)
	hexChars := strings.Split(os.Args[2], " ")
	packet := new(propellerhead.IbusPacket)
	packet.Src = hexChars[0]
	packet.Dest = hexChars[1]
	packet.Message = hexChars[2:len(hexChars)]
	fmt.Println("==> " + packet.AsString())
	bytes := packet.AsBytes()
	serialPort.Write(bytes)
}

