package propellerhead

type IbusPacketRouter struct {
	in chan *IbusPacket
}

func NewIbusPacketRouter(in chan *IbusPacket) (*IbusPacketRouter) {
	r := new(IbusPacketRouter)
	r.in = in
	return r
}

func (r *IbusPacketRouter) Listen() {
	go func() {
		for {
			p := <- r.in
			r.route(p)
		}
	}()
}

func (r *IbusPacketRouter) route(p *IbusPacket) {
	switch (p.Dest) {
	case IBUS_DEVICE_CDPLAYER:
		IbusDevices().CdPlayer.Handle(p)
	case IBUS_DEVICE_NAV_COMPUTER:
		IbusDevices().NavigationComputer.Handle(p)
	case IBUS_DEVICE_BROADCAST:
		IbusDevices().Broadcast.Handle(p)
	}

//	switch (p.Src) {
//
//	case DEVICE_RADIO:
//		fmt.Print("RADIO -> ")
//		switch(p.Dest) {
//		case DEVICE_CDPLAYER:
//			fmt.Print("CDPLAYER: ")
//			if (p.messageIs([]string{"01"})) {
//				fmt.Println("ping")
//			}
//		case DEVICE_BROADCAST:
//			fmt.Print("BROADCAST: ")
//			if (p.messageIs([]string{"02", "00"})) {
//				fmt.Println("pong")
//			}
//		}
//
//	case DEVICE_CDPLAYER:
//		fmt.Print("CDPLAYER -> ")
//		switch (p.Dest) {
//		case DEVICE_RADIO:
//			fmt.Print("RADIO ")
//			if (p.messageIs([]string{"02", "00"})) {
//				fmt.Println("pong")
//			}
//		}
//
//
//	case DEVICE_IKE:
//		fmt.Print("IKE -> ")
//		switch (p.Dest) {
//		case DEVICE_LIGHT_CONTROL:
//			fmt.Print("LCM: ")
//			fmt.Println("discarded")
//		}
//
//	case DEVICE_BOARD_MONITOR_BUTTONS:
//		fmt.Print("BMBUTTONS -> ")
//		switch (p.Dest) {
//		case DEVICE_RADIO:
//			fmt.Print("RADIO: ")
//			if (p.messageIs([]string{"01"})) {
//				fmt.Println("ping")
//			}
//			if (p.messageIs([]string{"48", "14"})) {
//				fmt.Println("reverse push")
//			}
//			if (p.messageIs([]string{"48", "54"})) {
//				fmt.Println("reverse hold")
//			}
//			if (p.messageIs([]string{"48", "94"})) {
//				fmt.Println("reverse release")
//			}
//			if (p.messageIs([]string{"48", "04"})) {
//				fmt.Println("tone push")
//			}
//			if (p.messageIs([]string{"48", "44"})) {
//				fmt.Println("tone hold")
//			}
//			if (p.messageIs([]string{"48", "84"})) {
//				fmt.Println("tone release")
//			}
//			if (p.messageIs([]string{"48", "00"})) {
//				fmt.Println("next push")
//			}
//			if (p.messageIs([]string{"48", "40"})) {
//				fmt.Println("next hold")
//			}
//			if (p.messageIs([]string{"48", "80"})) {
//				fmt.Println("next release")
//			}
//			if (p.messageIs([]string{"48", "10"})) {
//				fmt.Println("prev push")
//			}
//			if (p.messageIs([]string{"48", "50"})) {
//				fmt.Println("prev hold")
//			}
//			if (p.messageIs([]string{"48", "90"})) {
//				fmt.Println("prev release")
//			}

}
