package main

import (
	"propellerhead"
	"sync"
	"fmt"
	"time"
)


func main() {
	portWrite := propellerhead.OpenSerialPort("tty1")
	portRead := propellerhead.OpenSerialPort("tty2")
//	config := serial.RawOptions
//	config.FlowControl = serial.FLOWCONTROL_RTSCTS
//	config.BitRate = 9600

	wait := &sync.WaitGroup{}

	wait.Add(1)

	go func(){
		for {
			pkt := new(propellerhead.IbusPacket)
			pkt.Src = "68"
			pkt.Dest = "18"
			pkt.Message = []string{"01"}
			portWrite.Write(pkt.AsBytes())
			fmt.Println(pkt.AsBytes())
			time.Sleep(1 * time.Second)
		}
	}()

	wait.Add(1)
	go func(){
		for {
			byte := portRead.Read()
			fmt.Println(byte)
		}
	}()

	wait.Wait()

}
