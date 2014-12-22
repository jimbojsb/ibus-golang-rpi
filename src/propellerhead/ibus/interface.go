package ibus

import (
	"fmt"
	"propellerhead/audio"
	"github.com/johnlauer/serial"
)

type Interface struct {
	inboundPackets chan *Packet
	outboundPackets chan *Packet
	router Router
	parser Parser
}

func NewInterface(ac *audio.Controller) (*Interface) {
	iface := new(Interface)
	iface.inboundPackets = make(chan *Packet)
	iface.outboundPackets = make(chan *Packet)
	iface.router.in = iface.inboundPackets
	iface.router.out = iface.outboundPackets
	iface.router.audioController = ac
	return iface
}

func (i *Interface) Listen(ioDevicePath string) {

	config := &serial.Config{Name: ioDevicePath, Baud: 9600, RtsOn: true}
	serialPort, _ := serial.OpenPort(config)

	go i.router.listen()

	go func() {
		for {
			pkt := <- i.outboundPackets
			fmt.Printf("WRITING PACKET: %s\n", pkt.AsString())
			serialPort.Write(pkt.AsBytes())
		}
	}()


	for {
		char := make([]byte, 1)
		serialPort.Read(char)
		i.parser.Push(char[0])
		if (i.parser.HasPacket()) {
			pkt := i.parser.GetPacket();
			fmt.Printf("%+v\n", pkt)
			i.inboundPackets <- &pkt
		}
	}
}

func (i *Interface) GetOutboundChannel() (chan *Packet) {
	return i.outboundPackets
}
