package propellerhead

const EVENT_IBUS_NAV_KNOB_PUSH = "nav_knob_push"
const EVENT_IBUS_NAV_KNOB_HOLD = "nav_knob_hold"
const EVENT_IBUS_NAV_KNOB_RELEASE = "nav_knob_release"

type IbusNavigationComputer struct {
}

func NewIbusNavigationComputer() *IbusNavigationComputer {
	nav := new(IbusNavigationComputer)
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
