package ibus

const DEVICE_RADIO = "68"
const DEVICE_CDPLAYER = "18"
const DEVICE_BOARD_MONITOR_BUTTONS = "f0"
const DEVICE_IKE = "80"
const DEVICE_STEERING_WHEEL = "50"
const DEVICE_NAV_COMPUTER = "3b"
const DEVICE_PARK_DISTANCE = "60"
const DEVICE_LIGHT_CONTROL = "bf"
const DEVICE_NAV_LOCATION = "d0"
const DEVICE_BROADCAST = "ff"

var DeviceNames = map[string]string {
	DEVICE_RADIO: "RADIO",
	DEVICE_CDPLAYER: "CDPLAYER",
	DEVICE_BOARD_MONITOR_BUTTONS: "BMBUTTONS",
	DEVICE_IKE: "IKE",
	DEVICE_STEERING_WHEEL: "MFWHEEL",
	DEVICE_NAV_COMPUTER: "NAV",
	DEVICE_PARK_DISTANCE: "PDC",
	DEVICE_LIGHT_CONTROL: "LCM",
	DEVICE_BROADCAST: "BROADCAST",
}

func isKnownDevice(deviceId string) (bool) {
	for k := range DeviceNames {
		if (deviceId == k) {
			return true;
		}
	}
	return false
}
