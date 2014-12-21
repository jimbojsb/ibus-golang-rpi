package ibus

import (
	"encoding/hex"
	"strings"
)

type Packet struct {
	Src     string
	Dest    string
	Message []string
	Checksum string
}

func (pkt *Packet) messageIs(m []string) (bool) {
	if (len(pkt.Message) == len(m)) {
		for c := 0; c < len(pkt.Message); c++ {
			if (pkt.Message[c] != m[c]) {
				return false;
			}
		}
		return true;
	}
	return false
}

func (pkt *Packet) getLength() (string) {
	length := len(pkt.Message) + 2
	return getHexStringFromByte(byte(length))
}

func (pkt *Packet) AsString() (string) {
	return strings.Join(pkt.AsStringSlice(), " ")
}

func (pkt *Packet) AsStringSlice() ([]string) {
	var output []string
	output = append(output, pkt.Src)
	output = append(output, pkt.getLength())
	output = append(output, pkt.Dest)
	for _, el := range pkt.Message {
		output = append(output, el)
	}
	if (pkt.Checksum == "") {
		pkt.CaclulateAndSaveChecksum()
	}
	output = append(output, pkt.Checksum)
	return output
}

func (pkt *Packet) AsBytes() ([]byte) {
	var output []byte
	for _, el := range pkt.AsStringSlice() {
		output = append(output, getByteFromHexString(el))
	}
	return output
}

func getByteFromHexString(hexStr string) (byte) {
	byte, _ := hex.DecodeString(hexStr)
	return byte[0]
}

func getHexStringFromByte(b byte) (string) {
	return hex.EncodeToString([]byte{b})
}

func getHexStringSliceFromByteSlice(bytes []byte) ([]string) {
	output := make([]string, 0)
	for _, el := range bytes {
		output = append(output, getHexStringFromByte(el))
	}
	return output
}

func (pkt *Packet) AsDebugString() (string) {
	output := DeviceNames[pkt.Src] + " -> " + DeviceNames[pkt.Dest] + ": " + strings.Join(pkt.Message, " ")
	return output
}

func (pkt *Packet) CalculateChecksum() (string){
	var xor byte
	xor = xor ^ getByteFromHexString(pkt.Src)
	xor = xor ^ getByteFromHexString(pkt.getLength())
	xor = xor ^ getByteFromHexString(pkt.Dest)

	for _, el := range pkt.Message {
		xor = xor ^ getByteFromHexString(el)
	}
	return getHexStringFromByte(xor)
}

func (pkt *Packet) CaclulateAndSaveChecksum() {
	checksum := pkt.CalculateChecksum()
	pkt.Checksum = checksum
}

func (pkt *Packet) IsValid() (bool) {
	expectedChecksum := pkt.Checksum
	actualChecksum := pkt.CalculateChecksum()
	return expectedChecksum == actualChecksum
}
