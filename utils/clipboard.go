package utils

import (
	"syscall"
	"unsafe"
)

func SetClipboardText(text string) error {
	ret, _, _ := procOpenClipboard.Call(0)
	if ret == 0 {
		return syscall.GetLastError()
	}
	defer procCloseClipboard.Call()

	procEmptyClipboard.Call()

	hMem, _, _ := procGlobalAlloc.Call(0x0042, uintptr(len(text)*2+2))
	if hMem == 0 {
		return syscall.GetLastError()
	}

	lpData, _, _ := procGlobalLock.Call(hMem)
	if lpData == 0 {
		return syscall.GetLastError()
	}

	utf16Ptr, err := syscall.UTF16PtrFromString(text)
	if err != nil {
		return err
	}
	procLstrcpy.Call(lpData, uintptr(unsafe.Pointer(utf16Ptr)))
	procGlobalUnlock.Call(hMem)

	// CF_UNICODETEXT = 13
	ret, _, _ = procSetClipboardData.Call(13, hMem)
	if ret == 0 {
		return syscall.GetLastError()
	}

	return nil
}
