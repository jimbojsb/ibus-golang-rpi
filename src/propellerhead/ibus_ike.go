package propellerhead

type IbusIke struct {

}

func NewIbusIke() (*IbusIke) {
	i := new (IbusIke)
	return i
}

func (i *IbusIke) WriteText(text string) {
	pkt := new(IbusPacket)
	pkt.Src = "30"
	pkt.Dest = IBUS_DEVICE_IKE
}
