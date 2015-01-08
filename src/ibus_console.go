package main

import (
	"fmt"
	"bufio"
	"os"
	"propellerhead"
	"strings"
	"github.com/mikepb/serial"
	"time"
)

func main() {

	ttyPath := os.Args[1];
	fmt.Println("Writing packets to: " + ttyPath)
	fmt.Println("Packets are formatted all lower case hex, [src] [dest] [message...]")
	fmt.Println("Quote ascii strings for text conversion")



	config := serial.RawOptions
	config.FlowControl = serial.FLOWCONTROL_RTSCTS
	config.BitRate = 9600

	serialPort, _ := config.Open(ttyPath)
	serialPort.SetReadDeadline(time.Time{})

	ports, _ := serial.PortByName(ttyPath)
	fmt.Printf("%+v", ports)

	go func() {
		parser := propellerhead.NewIbusPacketParser()
		for {
			byte := make([]byte, 1)
			serialPort.Read(byte)
			parser.Push(byte[0])
			if (parser.HasPacket()) {
				pkt := parser.GetPacket();
				fmt.Println("\n<== " + pkt.AsString())
				fmt.Print("Enter IBUS packet: ")
			}
		}
	}()

	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print("Enter IBUS packet: ")
		text, _ := reader.ReadString('\n')
		text = strings.TrimSpace(text)
		hexChars := strings.Split(text, " ")
		packet := new(propellerhead.IbusPacket)
		packet.Src = hexChars[0]
		packet.Dest = hexChars[1]
		packet.Message = hexChars[2:len(hexChars)]
		fmt.Println("==> " + packet.AsString())
		bytes := packet.AsBytes()
		serialPort.Write(bytes)
	}
}

