package main

import (
	"fmt"
	"os"
	"propellerhead"
	"strings"
	"io/ioutil"
)

func main() {
	filepath := os.Args[1];

	hexChars := strings.Split(os.Args[2], " ")
	packet := new(propellerhead.IbusPacket)
	packet.Src = hexChars[0]
	packet.Dest = hexChars[1]
	packet.Message = hexChars[2:len(hexChars)]
	fmt.Println("==> " + packet.AsString())
	bytes := packet.AsBytes()
	ioutil.WriteFile(filepath, bytes, 0666)
}

