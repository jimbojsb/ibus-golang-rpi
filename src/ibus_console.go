package main

import (
	"fmt"
	"bufio"
	"os"
	"propellerhead"
	"strings"
)

func main() {

	ttyPath := os.Args[1];
	fmt.Println("Writing packets to: " + ttyPath)
	fmt.Println("Packets are formatted all lower case hex, [src] [dest] [message...]")
	fmt.Println("Quote ascii strings for text conversion")

	serialPort := propellerhead.OpenSerialPort(ttyPath)

	go func() {
		parser := propellerhead.NewIbusPacketParser()
		for {
			nextByte := serialPort.Read()
			parser.Push(nextByte)
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

