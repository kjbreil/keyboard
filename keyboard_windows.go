package keyboard

import (
	"fmt"
	"log"
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

// lazy loading the dll for now
var (
	userDLL     = windows.NewLazyDLL("user32.dll")
	keyEventDLL = userDLL.NewProc("keybd_event")

	setForegroundWindow = userDLL.NewProc("SetForegroundWindow")
	findWindowW         = userDLL.NewProc("FindWindowW")
	enumWindows         = userDLL.NewProc("EnumWindows")
	getWindowTextW      = userDLL.NewProc("GetWindowTextW")
)

// downKey is the down action on a key, this is how modifiers are added to a key
// press
func downKey(key rune) error {
	var vkey rune
	scanCode, ok := Scan[key]
	if ok {
		key = rune(scanCode.virtual)
		vkey = rune(scanCode.scan)
		log.Printf("Key: %s, Virtual: %v, Scan:%v\n", scanCode.name, key, vkey)

	} else {
		vkey = key + 0x80
	}
	_, _, err := keyEventDLL.Call(uintptr(key), uintptr(vkey), 0, 0)
	// error return needs to be corrected before here because its always filled
	// even if there is no error
	return err
}

// upKey the opposite of downkey
func upKey(key rune) error {
	var vkey rune
	scanCode, ok := Scan[key]
	if ok {
		key = rune(scanCode.virtual)
		vkey = rune(scanCode.scan)

	} else {
		vkey = key + 0x80
	}

	// defining keyUp to be more "verbose"
	var keyUp uintptr = 0x0002
	_, _, err := keyEventDLL.Call(uintptr(key), uintptr(vkey), keyUp, 0)
	return err
}

// SetForegroundWindow sets the window to the hwnd of the window, need to find
// that first :-)
func SetForegroundWindow(hwnd uintptr) bool {
	ret, _, _ := setForegroundWindow.Call(hwnd)

	return ret != 0
}

// FindWindow finds a window by name, needs to be exact
func FindWindow(win string) (ret uintptr, err error) {
	lpszWindow := syscall.StringToUTF16Ptr(win)

	r0, _, e1 := syscall.Syscall(findWindowW.Addr(), 2, uintptr(unsafe.Pointer(nil)), uintptr(unsafe.Pointer(lpszWindow)), 0)
	ret = uintptr(r0)
	if ret == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// Enumurating the windows is difficult. This is just a modification of
// this wonderful piece of code https://play.golang.org/p/YfGDtIuuBw
// the dll function needs a callback function to be passed to it
// https://docs.microsoft.com/en-us/windows/desktop/api/winuser/nf-winuser-enumwindows

// EnumWindows passes the functions to be run to the EnumWindow user32.dll
// function
func EnumWindows(enumFunc uintptr, lparam uintptr) (err error) {
	r1, _, e1 := syscall.Syscall(enumWindows.Addr(), 2, uintptr(enumFunc), uintptr(lparam), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// GetWindowText gets the window name text
func GetWindowText(hwnd syscall.Handle, str *uint16, maxCount int32) (len int32, err error) {
	r0, _, e1 := syscall.Syscall(getWindowTextW.Addr(), 3, uintptr(hwnd), uintptr(unsafe.Pointer(str)), uintptr(maxCount))
	len = int32(r0)
	if len == 0 {
		if e1 != 0 {
			err = error(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

// ListWindowNames lists all open window names
func ListWindowNames() error {
	cb := syscall.NewCallback(func(h syscall.Handle, p uintptr) uintptr {
		b := make([]uint16, 200)
		GetWindowText(h, &b[0], int32(len(b)))
		fmt.Println(syscall.UTF16ToString(b))
		return 1 // continue enumeration
	})
	EnumWindows(cb, 0)
	return nil
}
