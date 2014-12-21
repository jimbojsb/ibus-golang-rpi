package ibus

import (
	"fmt"
	"strconv"
)

type Parser struct {
	buffer    []byte
	packet    Packet
	hasPacket bool
}

func (p *Parser) Push(byte byte) {
	p.buffer = append(p.buffer, byte)
}

func (p *Parser) HasPacket() bool {
	p.parse()
	return p.hasPacket
}

func (p *Parser) GetPacket() Packet {
	p.hasPacket = false
	return p.packet
}

func (p *Parser) parse() {
	p.debug()
	p.hasPacket = false

	if (len(p.buffer)) < 5 {
		fmt.Println("packet buffer not long enough")
		return // all packets will be at least 5 bytes long
	}

	tmpLength, _ := strconv.ParseInt(getHexStringFromByte(p.buffer[1]), 16, 0)
	length := int(tmpLength)

	if (length == 0) {
		fmt.Println("packet length field cannot be 0")
		p.shiftBuffer()
		return
	}

	if (length < 3) {
		fmt.Println("packet length field not long enough (" + strconv.Itoa(length) + ")")
		p.shiftBuffer()
		return // not possible to have a length value less than 3 (dest + data + checksum)
	} else if (length > 100) {
		fmt.Println("packet length field too long (" + strconv.Itoa(length) + ")")
		p.shiftBuffer()
		return
	}


	if (!isKnownDevice(getHexStringFromByte(p.buffer[0]))) { // don't bother continuing if the source isn't a device we know
		fmt.Println("source byte is not a known device")
		p.shiftBuffer()
		return
	}

	if (len(p.buffer) < (2 + length)) {
		fmt.Println("Assuming length of " + strconv.Itoa(length) + ", we need more bytes")
		return // current byte slice is not long enough assuming length value at position 2
	} else {
		packet := new(Packet)
		packet.Src = getHexStringFromByte(p.buffer[0])
		packet.Dest = getHexStringFromByte(p.buffer[2])
		packet.Message = getHexStringSliceFromByteSlice(p.buffer[3:(3 + length - 2)])
		packet.Checksum = getHexStringFromByte(p.buffer[2 + length - 1])
		if (packet.IsValid()) {
			p.hasPacket = true
			p.packet = *packet
			p.buffer = p.buffer[2 + length : len(p.buffer)]
		} else {
			// we have enough bytes to theoretically be a valid packet, but either the checksum failed for unknown
			// reasons (unlikely), or buffer[0] is not actually the beggining of a packet. In this case, we will shift
			// off the first byte, as we know it is useless. This then means we are too short, but the new buffer[0]
			// might be a correct length byte, such that the next byte pushed onto the buffer will create a valid packet
			fmt.Println("packet validation error.")
			p.shiftBuffer()
		}
	}
}

func (p *Parser) shiftBuffer() {
	fmt.Println("Shifting buffer")
	p.buffer = p.buffer[1:len(p.buffer)]
	p.debug()
}

func (p *Parser) debug() {
	fmt.Println("")
	fmt.Println("")
	for _, el := range p.buffer {
		fmt.Print(getHexStringFromByte(el) + " ")
	}
	fmt.Println("")
}
