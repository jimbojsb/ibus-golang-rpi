package main

import (
	"fmt"
	"os"
	"propellerhead"
	"strings"
	"github.com/mikepb/serial"
	"time"
)

func main() {
	ttyPath := os.Args[1];

	config := serial.RawOptions
	config.FlowControl = serial.FLOWCONTROL_RTSCTS
	config.BitRate = 9600

	serialPort, _ := config.Open(ttyPath)
	serialPort.SetReadDeadline(time.Time{})

	hexChars := strings.Split(os.Args[2], " ")
	packet := new(propellerhead.IbusPacket)
	packet.Src = hexChars[0]
	packet.Dest = hexChars[1]
	packet.Message = hexChars[2:len(hexChars)]
	fmt.Println("==> " + packet.AsString())
	bytes := packet.AsBytes()
	serialPort.Write(bytes)
}

