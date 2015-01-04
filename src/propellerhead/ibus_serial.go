package propellerhead

import (
	"fmt"
	"github.com/johnlauer/serial"
)

type SerialInterface struct {
	inboundPackets chan *IbusPacket
	outboundPackets chan *IbusPacket
	parser *IbusPacketParser
	router *IbusPacketRouter
}

func NewSerialInterface() (*SerialInterface) {
	iface := new(SerialInterface)
	iface.inboundPackets = make(chan *IbusPacket, 32)
	iface.outboundPackets = make(chan *IbusPacket, 32)
	iface.parser = NewIbusPacketParser()
	iface.router = NewIbusPacketRouter(iface.inboundPackets)
	iface.router.Listen()
	return iface
}

func (i *SerialInterface) Write (pkt *IbusPacket) {
	i.outboundPackets <- pkt
}

func (i *SerialInterface) Listen(ioDevicePath string) {

	config := &serial.Config{Name: ioDevicePath, Baud: 9600, RtsOn: true}
	serialPort, _ := serial.OpenPort(config)

	go func() {
		for {
			pkt := <- i.outboundPackets
			fmt.Printf("WRITING PACKET: %s\n", pkt.AsString())
			serialPort.Write(pkt.AsBytes())
		}
	}()

	go func() {
		for {
			char := make([]byte, 1)
			serialPort.Read(char)
			i.parser.Push(char[0])
			if (i.parser.HasPacket()) {
				pkt := i.parser.GetPacket();
				fmt.Printf("%+v\n", pkt)
				i.inboundPackets <- pkt
			}
		}
	}()
}
