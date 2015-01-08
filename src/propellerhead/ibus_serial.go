package propellerhead

import (
	"github.com/mikepb/serial"
)

type IbusSerialInterface struct {
	inboundPackets chan *IbusPacket
	outboundPackets chan *IbusPacket
	parser *IbusPacketParser
	router *IbusPacketRouter
	LogOnly bool
}

func NewSerialInterface() (*IbusSerialInterface) {
	iface := new(IbusSerialInterface)
	iface.inboundPackets = make(chan *IbusPacket, 32)
	iface.outboundPackets = make(chan *IbusPacket, 32)
	iface.parser = NewIbusPacketParser()
	iface.router = NewIbusPacketRouter(iface.inboundPackets)
	iface.LogOnly = false
	iface.router.Listen()
	return iface
}

func (i *IbusSerialInterface) Write (pkt *IbusPacket) {
	i.outboundPackets <- pkt
}

func (i *IbusSerialInterface) Listen(ioDevicePath string) {

	config := serial.RawOptions
	config.FlowControl = serial.FLOWCONTROL_RTSCTS
	config.BitRate = 9600

	serialPort, _ := config.Open(ioDevicePath)
	defer serialPort.Close()

	if (!i.LogOnly) {
		go func() {
			for {
				pkt := <- i.outboundPackets
				Logger().Debug("sent packet " + pkt.AsString())
				serialPort.Write(pkt.AsBytes())
			}
		}()
	}

	go func() {
		for {
			char := make([]byte, 1)
			serialPort.Read(char)
			i.parser.Push(char[0])
			if (i.parser.HasPacket()) {
				pkt := i.parser.GetPacket();
				Logger().Debug("received packet " + pkt.AsString())
				if (!i.LogOnly) {
					i.inboundPackets <- pkt
				}
			}
		}
	}()
}
