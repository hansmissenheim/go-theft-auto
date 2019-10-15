// File keyboard implements functions to input key presses to the game.
// Code adapted from the snippet: https://github.com/golang/go/issues/31685.
package main

import (
	"syscall"
	"unsafe"
)

type keyboardInput struct {
	wVk         uint16
	wScan       uint16
	dwFlags     uint32
	time        uint32
	dwExtraInfo uint64
}

type input struct {
	inputType uint32
	ki        keyboardInput
	padding   uint64
}

// Scan codes for each key used
const (
	w = 0x11
	a = 0x1E
	s = 0x1F
	d = 0x20
)

var (
	user32        = syscall.NewLazyDLL("user32.dll")
	sendInputProc = user32.NewProc("SendInput")
	i             = input{inputType: 1}
)

// Presses a virtual key
func pressKey(key uint16) {
	i.ki.wScan = key
	i.ki.dwFlags = 0x0008
	sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		uintptr(unsafe.Sizeof(i)),
	)
}

// Releases a previously pressed virtual key
func releaseKey(key uint16) {
	i.ki.wScan = key
	i.ki.dwFlags = 0x0008 | 0x0002
	sendInputProc.Call(
		uintptr(1),
		uintptr(unsafe.Pointer(&i)),
		uintptr(unsafe.Sizeof(i)),
	)
}
