package main

import (
	"propellerhead/ibus"
	"strings"
	"io/ioutil"
	"propellerhead/prefs"
	"encoding/hex"
	"fmt"
)

func main() {
	p := new(ibus.Parser)

	fileContents, _ := ioutil.ReadFile(prefs.GetWorkingDir() + "/testcase")
	hexBytes := strings.Split(string(fileContents), " ")
	for _, el := range hexBytes {
		byte, _ := hex.DecodeString(el)
		p.Push(byte[0])
		if (p.HasPacket()) {
			packet := p.GetPacket()
			fmt.Println("--> Packet: " + packet.AsString())
		}
	}
}

