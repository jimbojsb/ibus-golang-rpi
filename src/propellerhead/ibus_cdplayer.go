package propellerhead

import (
	"time"
)

const EVENT_IBUS_CD_NEXT = "cd_next"
const EVENT_IBUS_CD_PREV = "cd_prev"
const EVENT_IBUS_CD_STOP = "cd_stop"
const EVENT_IBUS_CD_PLAY = "cd_play"

type CdPlayer struct {
	hasBeenPolled bool
}

func NewCdPlayer() *CdPlayer {
	cdp := new(CdPlayer)
	cdp.hasBeenPolled = false
	cdp.announce()
	return cdp
}

func (cdp *CdPlayer) Handle(p *IbusPacket) {
	switch p.Src {
	case IBUS_DEVICE_RADIO:
		if p.messageIs([]string{"01"}) {
			cdp.pong()
		}
		if p.messageIs([]string{"38", "0a", "00"}) {
			Emitter().Emit(EVENT_IBUS_CD_NEXT)
		}
		if p.messageIs([]string{"38", "0a", "01"}) {
			Emitter().Emit(EVENT_IBUS_CD_PREV)
		}
		if p.messageIs([]string{"38", "01", "00"}) {
			Emitter().Emit(EVENT_IBUS_CD_STOP)
		}
		if p.messageIs([]string{"38", "03", "00"}) {
			Emitter().Emit(EVENT_IBUS_CD_PLAY)
		}
	}
}

func (cdp *CdPlayer) announce() {
	go func() {
		pkt := new(IbusPacket)
		pkt.Src = IBUS_DEVICE_CDPLAYER
		pkt.Dest = IBUS_DEVICE_BROADCAST
		pkt.Message = []string{"02", "01"}
		for {
			if cdp.hasBeenPolled {
				return
			}
			IbusDevices().SerialInterface.Write(pkt)
			time.Sleep(30 * time.Second)
		}
	}()
}

func (cdp *CdPlayer) pong() {
	cdp.hasBeenPolled = true
	pkt := new(IbusPacket)
	pkt.Src = IBUS_DEVICE_CDPLAYER
	pkt.Dest = IBUS_DEVICE_BROADCAST
	pkt.Message = []string{"02", "00"}
	IbusDevices().SerialInterface.Write(pkt)
}


