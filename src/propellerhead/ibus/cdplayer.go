package ibus

import "propellerhead/audio"


type CdPlayer struct {
	out chan *Packet
	audioController *audio.Controller
}

func (cdp *CdPlayer) handle(p *Packet) {
	switch (p.Src) {
	case DEVICE_RADIO:
		if (p.messageIs([]string{"01"})) {
			cdp.out <- cdp.pong(p)
		}
		if (p.messageIs([]string{"38", "0a", "00"})) {
			cdp.nextTrack()
		}
		if (p.messageIs([]string{"38", "0a", "01"})) {
			cdp.prevTrack()
		}
		if (p.messageIs([]string{"38", "01", "00"})) {
			cdp.stop()
		}
		if (p.messageIs([]string{"38", "03", "00"})) {
			cdp.play()
		}
	}
}

func (cdp *CdPlayer) pong(p *Packet) (*Packet) {
	pkt := new(Packet)
	pkt.Src = DEVICE_CDPLAYER
	pkt.Dest = DEVICE_BROADCAST
	pkt.Message = []string{"02", "00"}
	return pkt
}

func (cdp *CdPlayer) nextTrack() {
	cdp.audioController.Next()
}

func (cdp *CdPlayer) prevTrack() {
	cdp.audioController.Prev()
}

func (cdp *CdPlayer) play() {
	cdp.audioController.Play()
}

func (cdp *CdPlayer) stop() {
	cdp.audioController.Pause()
}
