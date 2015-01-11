package propellerhead

type IbusRadio struct {

}

func NewIbusRadio() *IbusRadio {
	rad := new(IbusRadio)
	return rad
}

func (cdp *IbusRadio) Handle(p *IbusPacket) {
	switch p.Src {
	case IBUS_DEVICE_MFW:
		if p.messageIs([]string{"3b", "01"}) {
			Emitter().Emit(EVENT_IBUS_MFW_NEXT_PUSH)
			Logger().Info("MFW next push")
		}
		if p.messageIs([]string{"3b", "21"}) {
			Emitter().Emit(EVENT_IBUS_MFW_NEXT_RELEASE)
			Logger().Info("MFW next release")
		}
		if p.messageIs([]string{"3b", "08"}) {
			Emitter().Emit(EVENT_IBUS_MFW_PREV_PUSH)
			Logger().Info("MFW prev push")
		}
		if p.messageIs([]string{"3b", "28"}) {
			Emitter().Emit(EVENT_IBUS_MFW_PREV_RELEASE)
			Logger().Info("MFW prev release")
		}
	}
}
