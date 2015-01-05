package propellerhead

const EVENT_IBUS_MENU_PRESS = "menu_press"
const EVENT_IBUS_MENU_HOLD = "menu_hold"
const EVENT_IBUS_MENU_RELEASE = "menu_release"

type IbusBroadcast struct {
}

func NewIbusBroadcast() *IbusBroadcast {
	bcast := new(IbusBroadcast)
	return bcast
}

func (bcast *IbusBroadcast) Handle(p *IbusPacket) {
	switch p.Src {
	case IBUS_DEVICE_BOARD_MONITOR_BUTTONS:
		if (p.messageIs([]string{"48", "34"})) {
			Logger().Info("menu press")
			Emitter().Emit(EVENT_IBUS_MENU_PRESS)
		}
		if (p.messageIs([]string{"48", "74"})) {
			Logger().Info("menu hold")
			Emitter().Emit(EVENT_IBUS_MENU_HOLD)
		}
		if (p.messageIs([]string{"48", "b4"})) {
			Logger().Info("menu release")
			Emitter().Emit(EVENT_IBUS_MENU_RELEASE)
		}
	}
}
