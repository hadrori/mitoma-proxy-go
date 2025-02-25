package utils

import (
	"syscall"
)

const (
	KEYEVENTF_EXTENDEDKEY = 0x0001
	KEYEVENTF_KEYUP       = 0x0002
)

var (
	user32   = syscall.MustLoadDLL("user32.dll")
	kernel32 = syscall.MustLoadDLL("kernel32.dll")

	procKeybd_Event      = user32.MustFindProc("keybd_event")
	procOpenClipboard    = user32.MustFindProc("OpenClipboard")
	procEmptyClipboard   = user32.MustFindProc("EmptyClipboard")
	procSetClipboardData = user32.MustFindProc("SetClipboardData")
	procCloseClipboard   = user32.MustFindProc("CloseClipboard")

	procGlobalAlloc  = kernel32.MustFindProc("GlobalAlloc")
	procGlobalLock   = kernel32.MustFindProc("GlobalLock")
	procGlobalUnlock = kernel32.MustFindProc("GlobalUnlock")
	procLstrcpy      = kernel32.MustFindProc("lstrcpyW")
)
