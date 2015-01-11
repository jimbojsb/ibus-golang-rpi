package propellerhead

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

func (i *IbusSerialInterface) Listen(ttyPath string) {

	serialPort := OpenSerialPort(ttyPath)

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
			nextByte := serialPort.Read()
			i.parser.Push(nextByte)
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
