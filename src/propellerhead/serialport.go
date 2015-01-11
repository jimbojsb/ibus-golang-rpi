package propellerhead

import (
	linux_serial "github.com/mikepb/serial"
	mac_serial "github.com/johnlauer/serial"
	"time"
	"io"
	"runtime"
)

type SerialPort struct {
	macPort *io.ReadWriteCloser
	linuxPort *linux_serial.Port
}

func OpenSerialPort(ttyPath string) (*SerialPort) {
	s := new(SerialPort)
	if (runtime.GOOS == "darwin") {
		c := &mac_serial.Config{Name: ttyPath, Baud: 9600, RtsOn: true}
		port, _ := mac_serial.OpenPort(c)
		s.macPort = &port
	} else {
		config := linux_serial.RawOptions
		config.FlowControl = linux_serial.FLOWCONTROL_RTSCTS
		config.BitRate = 9600

		port, _ := config.Open(ttyPath)
		port.SetReadDeadline(time.Time{})
		s.linuxPort = port
	}
	return s
}

func (s *SerialPort) Read() ([]byte) {
	b := make([]byte, 1)
	if (runtime.GOOS == "darwin") {
		port := *s.macPort
		port.Read(b)
	} else {
		s.linuxPort.Read(b)
	}
	return b
}

func (s *SerialPort) Write(data []byte) {
	if (runtime.GOOS == "darwin") {
		port := *s.macPort
		port.Write(data)
	} else {
		s.linuxPort.Write(data)
	}
}
