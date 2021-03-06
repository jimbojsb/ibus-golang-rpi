package propellerhead

import (
	"unicode/utf8"
	"strings"
)

const EVENT_IBUS_NAV_KNOB_PUSH = "nav_knob_push"
const EVENT_IBUS_NAV_KNOB_HOLD = "nav_knob_hold"
const EVENT_IBUS_NAV_KNOB_RELEASE = "nav_knob_release"

const IBUS_NAV_TEXT_AREA_0 = "00"
const IBUS_NAV_TEXT_AREA_1 = "01"
const IBUS_NAV_TEXT_AREA_2 = "02"
const IBUS_NAV_TEXT_AREA_3 = "03"
const IBUS_NAV_TEXT_AREA_4 = "04"
const IBUS_NAV_TEXT_AREA_5 = "05"
const IBUS_NAV_TEXT_AREA_6 = "06"
const IBUS_NAV_TEXT_AREA_7 = "07"




type IbusNavigationComputer struct {
}

func NewIbusNavigationComputer() *IbusNavigationComputer {
	nav := new(IbusNavigationComputer)
	nav.bindEvents()
	return nav
}

func (nav *IbusNavigationComputer) Handle(p *IbusPacket) {
	switch p.Src {
	case IBUS_DEVICE_BOARD_MONITOR_BUTTONS:
		if (p.messageIs([]string{"48", "05"})) {
			Logger().Info("nav knob push")
			Emitter().Emit(EVENT_IBUS_NAV_KNOB_PUSH)
		}
		if (p.messageIs([]string{"48", "45"})) {
			Logger().Info("nav knob hold")
			Emitter().Emit(EVENT_IBUS_NAV_KNOB_HOLD)
		}
		if (p.messageIs([]string{"48", "85"})) {
			Logger().Info("nav knob release")
			Emitter().Emit(EVENT_IBUS_NAV_KNOB_RELEASE)
		}
	}
}

func (nav *IbusNavigationComputer) bindEvents() {
	Emitter().On(EVENT_AUDIO_SOURCE_CHANGED, func(source string) {
		nav.WriteTextArea(IBUS_NAV_TEXT_AREA_0, strings.Title(source))
	})
}

func (nav *IbusNavigationComputer) WriteTextArea(area string, text string) {
	pkt := new(IbusPacket)
	var hexChars []string
	textLength := utf8.RuneCountInString(text)

	if (area == IBUS_NAV_TEXT_AREA_0) {
		len := 11
		if (textLength < 11) {
			len = textLength
		}
		text = text[0:len]
		hexChars = []string{"23", "62", "30"}
	} else {
		hexChars = []string{"a5", "62", "01", area}
		if (area == IBUS_NAV_TEXT_AREA_1 || area == IBUS_NAV_TEXT_AREA_2 || area == IBUS_NAV_TEXT_AREA_3 || area == IBUS_NAV_TEXT_AREA_4) {
			text = text[0:5]
		} else if (area == IBUS_NAV_TEXT_AREA_5) {
			text = text[0:7]
		} else {
			text = text[0:20]
		}
	}

	hexChars = append(hexChars, stringAsHexStringSlice(text)...)
	pkt.Src = IBUS_DEVICE_RADIO
	pkt.Dest = IBUS_DEVICE_NAV_COMPUTER
	pkt.Message = hexChars
	IbusDevices().SerialInterface.Write(pkt)
}
