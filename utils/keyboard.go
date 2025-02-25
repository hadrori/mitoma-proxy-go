package utils

import (
	"mitoma-proxy-go/constants"
	"time"
)

func SendKeyPress(keyCode byte) {
	procKeybd_Event.Call(uintptr(keyCode),
		0,
		KEYEVENTF_EXTENDEDKEY,
		0)

	time.Sleep(1 * time.Millisecond)

	procKeybd_Event.Call(uintptr(keyCode),
		0,
		KEYEVENTF_EXTENDEDKEY|KEYEVENTF_KEYUP,
		0)
}

func SendCtrlKeyCombo(keyCode byte) {
	procKeybd_Event.Call(uintptr(constants.VK_CONTROL),
		0,
		KEYEVENTF_EXTENDEDKEY,
		0)

	time.Sleep(3 * time.Millisecond)

	SendKeyPress(keyCode)

	time.Sleep(1 * time.Millisecond)

	procKeybd_Event.Call(uintptr(constants.VK_CONTROL),
		0,
		KEYEVENTF_EXTENDEDKEY|KEYEVENTF_KEYUP,
		0)
}

func SendCtrlL() {
	SendCtrlKeyCombo(constants.VK_L)
}

func SendCtrlV() {
	SendCtrlKeyCombo(constants.VK_V)
}

func SendEnter() {
	SendKeyPress(constants.VK_RETURN)
}
