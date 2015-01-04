package propellerhead

const IBUS_DEVICE_RADIO = "68"
const IBUS_DEVICE_CDPLAYER = "18"
const IBUS_DEVICE_BOARD_MONITOR_BUTTONS = "f0"
const IBUS_DEVICE_IKE = "80"
const IBUS_DEVICE_STEERING_WHEEL = "50"
const IBUS_DEVICE_NAV_COMPUTER = "3b"
const IBUS_DEVICE_PARK_DISTANCE = "60"
const IBUS_DEVICE_LIGHT_CONTROL = "bf"
const IBUS_DEVICE_NAV_LOCATION = "d0"
const IBUS_DEVICE_BROADCAST = "ff"

type IbusDeviceRegistry struct {
	CdPlayer *CdPlayer
	SerialInterface *SerialInterface
}

var IbusDeviceNames = map[string]string {
	IBUS_DEVICE_RADIO: "RADIO",
	IBUS_DEVICE_CDPLAYER: "CDPLAYER",
	IBUS_DEVICE_BOARD_MONITOR_BUTTONS: "BMBUTTONS",
	IBUS_DEVICE_IKE: "IKE",
	IBUS_DEVICE_STEERING_WHEEL: "MFWHEEL",
	IBUS_DEVICE_NAV_COMPUTER: "NAV",
	IBUS_DEVICE_PARK_DISTANCE: "PDC",
	IBUS_DEVICE_LIGHT_CONTROL: "LCM",
	IBUS_DEVICE_BROADCAST: "BROADCAST",
}

var ibusDeviceregistry *IbusDeviceRegistry

func IbusDevices() (*IbusDeviceRegistry) {
	if (ibusDeviceregistry == nil) {
		ibusDeviceregistry = new(IbusDeviceRegistry)
		ibusDeviceregistry.CdPlayer = NewCdPlayer()
		ibusDeviceregistry.SerialInterface = NewSerialInterface()
	}
	return ibusDeviceregistry
}

func IsKnownIbusDevice(deviceId string) (bool) {
	for k := range IbusDeviceNames {
		if (deviceId == k) {
			return true;
		}
	}
	return false
}

