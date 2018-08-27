// MACHINE GENERATED BY 'go generate' COMMAND; DO NOT EDIT

package windriver

import (
	"syscall"
	"unsafe"

	"golang.org/x/sys/windows"
)

var _ unsafe.Pointer

// Do the interface allocations only once for common
// Errno values.
const (
	errnoERROR_IO_PENDING = 997
)

var (
	errERROR_IO_PENDING error = syscall.Errno(errnoERROR_IO_PENDING)
)

// errnoErr returns common boxed Errno values, to prevent
// allocations at runtime.
func errnoErr(e syscall.Errno) error {
	switch e {
	case 0:
		return nil
	case errnoERROR_IO_PENDING:
		return errERROR_IO_PENDING
	}
	// TODO: add more here, after collecting data on the common
	// error values see on Windows. (perhaps when running
	// all.bat?)
	return e
}

var (
	moduser32   = windows.NewLazySystemDLL("user32.dll")
	modkernel32 = windows.NewLazySystemDLL("kernel32.dll")
	modmsimg32  = windows.NewLazySystemDLL("msimg32.dll")
	modgdi32    = windows.NewLazySystemDLL("gdi32.dll")
	modshcore   = windows.NewLazySystemDLL("shcore.dll")

	procGetDC                      = moduser32.NewProc("GetDC")
	procReleaseDC                  = moduser32.NewProc("ReleaseDC")
	procSendMessageW               = moduser32.NewProc("SendMessageW")
	procCreateWindowExW            = moduser32.NewProc("CreateWindowExW")
	procDefWindowProcW             = moduser32.NewProc("DefWindowProcW")
	procDestroyWindow              = moduser32.NewProc("DestroyWindow")
	procDispatchMessageW           = moduser32.NewProc("DispatchMessageW")
	procGetClientRect              = moduser32.NewProc("GetClientRect")
	procGetWindowRect              = moduser32.NewProc("GetWindowRect")
	procGetKeyboardLayout          = moduser32.NewProc("GetKeyboardLayout")
	procGetKeyboardState           = moduser32.NewProc("GetKeyboardState")
	procGetKeyState                = moduser32.NewProc("GetKeyState")
	procGetMessageW                = moduser32.NewProc("GetMessageW")
	procLoadCursorA                = moduser32.NewProc("LoadCursorA")
	procLoadIconW                  = moduser32.NewProc("LoadIconW")
	procSetCursor                  = moduser32.NewProc("SetCursor")
	procShowCursor                 = moduser32.NewProc("ShowCursor")
	procMoveWindow                 = moduser32.NewProc("MoveWindow")
	procBringWindowToTop           = moduser32.NewProc("BringWindowToTop")
	procPostMessageW               = moduser32.NewProc("PostMessageW")
	procPostQuitMessage            = moduser32.NewProc("PostQuitMessage")
	procRegisterClassW             = moduser32.NewProc("RegisterClassW")
	procShowWindow                 = moduser32.NewProc("ShowWindow")
	procSetActiveWindow            = moduser32.NewProc("SetActiveWindow")
	procSetFocus                   = moduser32.NewProc("SetFocus")
	procSetForegroundWindow        = moduser32.NewProc("SetForegroundWindow")
	procIsIconic                   = moduser32.NewProc("IsIconic")
	procIsWindowVisible            = moduser32.NewProc("IsWindowVisible")
	procScreenToClient             = moduser32.NewProc("ScreenToClient")
	procToUnicodeEx                = moduser32.NewProc("ToUnicodeEx")
	procTranslateMessage           = moduser32.NewProc("TranslateMessage")
	procOpenClipboard              = moduser32.NewProc("OpenClipboard")
	procCloseClipboard             = moduser32.NewProc("CloseClipboard")
	procEmptyClipboard             = moduser32.NewProc("EmptyClipboard")
	procSetClipboardData           = moduser32.NewProc("SetClipboardData")
	procGetClipboardData           = moduser32.NewProc("GetClipboardData")
	procIsClipboardFormatAvailable = moduser32.NewProc("IsClipboardFormatAvailable")
	procGlobalLock                 = modkernel32.NewProc("GlobalLock")
	procGlobalUnlock               = modkernel32.NewProc("GlobalUnlock")
	procGlobalAlloc                = modkernel32.NewProc("GlobalAlloc")
	procGlobalFree                 = modkernel32.NewProc("GlobalFree")
	procRtlCopyMemory              = modkernel32.NewProc("RtlCopyMemory")
	procAlphaBlend                 = modmsimg32.NewProc("AlphaBlend")
	procBitBlt                     = modgdi32.NewProc("BitBlt")
	procCreateCompatibleBitmap     = modgdi32.NewProc("CreateCompatibleBitmap")
	procCreateCompatibleDC         = modgdi32.NewProc("CreateCompatibleDC")
	procCreateDIBSection           = modgdi32.NewProc("CreateDIBSection")
	procCreateSolidBrush           = modgdi32.NewProc("CreateSolidBrush")
	procDeleteDC                   = modgdi32.NewProc("DeleteDC")
	procDeleteObject               = modgdi32.NewProc("DeleteObject")
	procFillRect                   = moduser32.NewProc("FillRect")
	procModifyWorldTransform       = modgdi32.NewProc("ModifyWorldTransform")
	procSelectObject               = modgdi32.NewProc("SelectObject")
	procSetGraphicsMode            = modgdi32.NewProc("SetGraphicsMode")
	procSetWorldTransform          = modgdi32.NewProc("SetWorldTransform")
	procStretchBlt                 = modgdi32.NewProc("StretchBlt")
	procGetDeviceCaps              = modgdi32.NewProc("GetDeviceCaps")
	procSetProcessDpiAwareness     = modshcore.NewProc("SetProcessDpiAwareness")
	procGetDpiForWindow            = moduser32.NewProc("GetDpiForWindow")
	procEnumDisplayDevicesA        = moduser32.NewProc("EnumDisplayDevicesA")
)

func _GetDC(hwnd syscall.Handle) (dc syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procGetDC.Addr(), 1, uintptr(hwnd), 0, 0)
	dc = syscall.Handle(r0)
	if dc == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _ReleaseDC(hwnd syscall.Handle, dc syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procReleaseDC.Addr(), 2, uintptr(hwnd), uintptr(dc), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _sendMessage(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	r0, _, _ := syscall.Syscall6(procSendMessageW.Addr(), 4, uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam), 0, 0)
	lResult = uintptr(r0)
	return
}

func _CreateWindowEx(exstyle uint32, className *uint16, windowText *uint16, style uint32, x int32, y int32, width int32, height int32, parent syscall.Handle, menu syscall.Handle, hInstance syscall.Handle, lpParam uintptr) (hwnd syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall12(procCreateWindowExW.Addr(), 12, uintptr(exstyle), uintptr(unsafe.Pointer(className)), uintptr(unsafe.Pointer(windowText)), uintptr(style), uintptr(x), uintptr(y), uintptr(width), uintptr(height), uintptr(parent), uintptr(menu), uintptr(hInstance), uintptr(lpParam))
	hwnd = syscall.Handle(r0)
	if hwnd == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _DefWindowProc(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) (lResult uintptr) {
	r0, _, _ := syscall.Syscall6(procDefWindowProcW.Addr(), 4, uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam), 0, 0)
	lResult = uintptr(r0)
	return
}

func _DestroyWindow(hwnd syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDestroyWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _DispatchMessage(msg *_MSG) (ret int32) {
	r0, _, _ := syscall.Syscall(procDispatchMessageW.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	ret = int32(r0)
	return
}

func _GetClientRect(hwnd syscall.Handle, rect *_RECT) (err error) {
	r1, _, e1 := syscall.Syscall(procGetClientRect.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(rect)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetWindowRect(hwnd syscall.Handle, rect *_RECT) (err error) {
	r1, _, e1 := syscall.Syscall(procGetWindowRect.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(rect)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetKeyboardLayout(threadID uint32) (locale syscall.Handle) {
	r0, _, _ := syscall.Syscall(procGetKeyboardLayout.Addr(), 1, uintptr(threadID), 0, 0)
	locale = syscall.Handle(r0)
	return
}

func _GetKeyboardState(lpKeyState *byte) (err error) {
	r1, _, e1 := syscall.Syscall(procGetKeyboardState.Addr(), 1, uintptr(unsafe.Pointer(lpKeyState)), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetKeyState(virtkey int32) (keystatus int16) {
	r0, _, _ := syscall.Syscall(procGetKeyState.Addr(), 1, uintptr(virtkey), 0, 0)
	keystatus = int16(r0)
	return
}

func _GetMessage(msg *_MSG, hwnd syscall.Handle, msgfiltermin uint32, msgfiltermax uint32) (ret int32, err error) {
	r0, _, e1 := syscall.Syscall6(procGetMessageW.Addr(), 4, uintptr(unsafe.Pointer(msg)), uintptr(hwnd), uintptr(msgfiltermin), uintptr(msgfiltermax), 0, 0)
	ret = int32(r0)
	if ret == -1 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _LoadCursor(hInstance syscall.Handle, cursorName uintptr) (cursor syscall.Handle) {
	r0, _, _ := syscall.Syscall(procLoadCursorA.Addr(), 2, uintptr(hInstance), uintptr(cursorName), 0)
	cursor = syscall.Handle(r0)
	return
}

func _LoadIcon(hInstance syscall.Handle, iconName uintptr) (icon syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procLoadIconW.Addr(), 2, uintptr(hInstance), uintptr(iconName), 0)
	icon = syscall.Handle(r0)
	if icon == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _SetCursor(hinst syscall.Handle) (prevcurs syscall.Handle) {
	r0, _, _ := syscall.Syscall(procSetCursor.Addr(), 1, uintptr(hinst), 0, 0)
	prevcurs = syscall.Handle(r0)
	return
}

func _ShowCursor(show bool) (dispcnt int) {
	var _p0 uint32
	if show {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r0, _, _ := syscall.Syscall(procShowCursor.Addr(), 1, uintptr(_p0), 0, 0)
	dispcnt = int(r0)
	return
}

func _MoveWindow(hwnd syscall.Handle, x int32, y int32, w int32, h int32, repaint bool) (err error) {
	var _p0 uint32
	if repaint {
		_p0 = 1
	} else {
		_p0 = 0
	}
	r1, _, e1 := syscall.Syscall6(procMoveWindow.Addr(), 6, uintptr(hwnd), uintptr(x), uintptr(y), uintptr(w), uintptr(h), uintptr(_p0))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _BringWindowToTop(hwnd syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procBringWindowToTop.Addr(), 1, uintptr(hwnd), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _PostMessage(hwnd syscall.Handle, uMsg uint32, wParam uintptr, lParam uintptr) (lResult bool) {
	r0, _, _ := syscall.Syscall6(procPostMessageW.Addr(), 4, uintptr(hwnd), uintptr(uMsg), uintptr(wParam), uintptr(lParam), 0, 0)
	lResult = r0 != 0
	return
}

func _PostQuitMessage(exitCode int32) {
	syscall.Syscall(procPostQuitMessage.Addr(), 1, uintptr(exitCode), 0, 0)
	return
}

func _RegisterClass(wc *_WNDCLASS) (atom uint16, err error) {
	r0, _, e1 := syscall.Syscall(procRegisterClassW.Addr(), 1, uintptr(unsafe.Pointer(wc)), 0, 0)
	atom = uint16(r0)
	if atom == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _ShowWindow(hwnd syscall.Handle, cmdshow int32) (wasvisible bool) {
	r0, _, _ := syscall.Syscall(procShowWindow.Addr(), 2, uintptr(hwnd), uintptr(cmdshow), 0)
	wasvisible = r0 != 0
	return
}

func _SetActiveWindow(hwnd syscall.Handle) (prev syscall.Handle) {
	r0, _, _ := syscall.Syscall(procSetActiveWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	prev = syscall.Handle(r0)
	return
}

func _SetFocus(hwnd syscall.Handle) (prev syscall.Handle) {
	r0, _, _ := syscall.Syscall(procSetFocus.Addr(), 1, uintptr(hwnd), 0, 0)
	prev = syscall.Handle(r0)
	return
}

func _SetForegroundWindow(hwnd syscall.Handle) (ok bool) {
	r0, _, _ := syscall.Syscall(procSetForegroundWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	ok = r0 != 0
	return
}

func _IsIconic(hwnd syscall.Handle) (iconic bool) {
	r0, _, _ := syscall.Syscall(procIsIconic.Addr(), 1, uintptr(hwnd), 0, 0)
	iconic = r0 != 0
	return
}

func _IsWindowVisible(hwnd syscall.Handle) (vis bool) {
	r0, _, _ := syscall.Syscall(procIsWindowVisible.Addr(), 1, uintptr(hwnd), 0, 0)
	vis = r0 != 0
	return
}

func _ScreenToClient(hwnd syscall.Handle, lpPoint *_POINT) (ok bool) {
	r0, _, _ := syscall.Syscall(procScreenToClient.Addr(), 2, uintptr(hwnd), uintptr(unsafe.Pointer(lpPoint)), 0)
	ok = r0 != 0
	return
}

func _ToUnicodeEx(wVirtKey uint32, wScanCode uint32, lpKeyState *byte, pwszBuff *uint16, cchBuff int32, wFlags uint32, dwhkl syscall.Handle) (ret int32) {
	r0, _, _ := syscall.Syscall9(procToUnicodeEx.Addr(), 7, uintptr(wVirtKey), uintptr(wScanCode), uintptr(unsafe.Pointer(lpKeyState)), uintptr(unsafe.Pointer(pwszBuff)), uintptr(cchBuff), uintptr(wFlags), uintptr(dwhkl), 0, 0)
	ret = int32(r0)
	return
}

func _TranslateMessage(msg *_MSG) (done bool) {
	r0, _, _ := syscall.Syscall(procTranslateMessage.Addr(), 1, uintptr(unsafe.Pointer(msg)), 0, 0)
	done = r0 != 0
	return
}

func _OpenClipboard(hwnd syscall.Handle) (opened bool) {
	r0, _, _ := syscall.Syscall(procOpenClipboard.Addr(), 1, uintptr(hwnd), 0, 0)
	opened = r0 != 0
	return
}

func _CloseClipboard() (closed bool) {
	r0, _, _ := syscall.Syscall(procCloseClipboard.Addr(), 0, 0, 0, 0)
	closed = r0 != 0
	return
}

func _EmptyClipboard() (empty bool) {
	r0, _, _ := syscall.Syscall(procEmptyClipboard.Addr(), 0, 0, 0, 0)
	empty = r0 != 0
	return
}

func _SetClipboardData(uFormat uint32, hMem syscall.Handle) (hRes syscall.Handle) {
	r0, _, _ := syscall.Syscall(procSetClipboardData.Addr(), 2, uintptr(uFormat), uintptr(hMem), 0)
	hRes = syscall.Handle(r0)
	return
}

func _GetClipboardData(uFormat uint32) (hMem syscall.Handle) {
	r0, _, _ := syscall.Syscall(procGetClipboardData.Addr(), 1, uintptr(uFormat), 0, 0)
	hMem = syscall.Handle(r0)
	return
}

func _IsClipboardFormatAvailable(uFormat uint32) (avail bool) {
	r0, _, _ := syscall.Syscall(procIsClipboardFormatAvailable.Addr(), 1, uintptr(uFormat), 0, 0)
	avail = r0 != 0
	return
}

func _GlobalLock(hMem syscall.Handle) (data *uint16) {
	r0, _, _ := syscall.Syscall(procGlobalLock.Addr(), 1, uintptr(hMem), 0, 0)
	data = (*uint16)(unsafe.Pointer(r0))
	return
}

func _GlobalUnlock(hMem syscall.Handle) (unlocked bool) {
	r0, _, _ := syscall.Syscall(procGlobalUnlock.Addr(), 1, uintptr(hMem), 0, 0)
	unlocked = r0 != 0
	return
}

func _GlobalAlloc(uFlags uint32, size uintptr) (hMem syscall.Handle) {
	r0, _, _ := syscall.Syscall(procGlobalAlloc.Addr(), 2, uintptr(uFlags), uintptr(size), 0)
	hMem = syscall.Handle(r0)
	return
}

func _GlobalFree(hMem syscall.Handle) {
	syscall.Syscall(procGlobalFree.Addr(), 1, uintptr(hMem), 0, 0)
	return
}

func _CopyMemory(dest uintptr, src uintptr, sz uintptr) {
	syscall.Syscall(procRtlCopyMemory.Addr(), 3, uintptr(dest), uintptr(src), uintptr(sz))
	return
}

func _AlphaBlend(dcdest syscall.Handle, xoriginDest int32, yoriginDest int32, wDest int32, hDest int32, dcsrc syscall.Handle, xoriginSrc int32, yoriginSrc int32, wsrc int32, hsrc int32, ftn uintptr) (err error) {
	r1, _, e1 := syscall.Syscall12(procAlphaBlend.Addr(), 11, uintptr(dcdest), uintptr(xoriginDest), uintptr(yoriginDest), uintptr(wDest), uintptr(hDest), uintptr(dcsrc), uintptr(xoriginSrc), uintptr(yoriginSrc), uintptr(wsrc), uintptr(hsrc), uintptr(ftn), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _BitBlt(dcdest syscall.Handle, xdest int32, ydest int32, width int32, height int32, dcsrc syscall.Handle, xsrc int32, ysrc int32, rop uint32) (err error) {
	r1, _, e1 := syscall.Syscall9(procBitBlt.Addr(), 9, uintptr(dcdest), uintptr(xdest), uintptr(ydest), uintptr(width), uintptr(height), uintptr(dcsrc), uintptr(xsrc), uintptr(ysrc), uintptr(rop))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateCompatibleBitmap(dc syscall.Handle, width int32, height int32) (bitmap syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleBitmap.Addr(), 3, uintptr(dc), uintptr(width), uintptr(height))
	bitmap = syscall.Handle(r0)
	if bitmap == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateCompatibleDC(dc syscall.Handle) (newdc syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateCompatibleDC.Addr(), 1, uintptr(dc), 0, 0)
	newdc = syscall.Handle(r0)
	if newdc == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateDIBSection(dc syscall.Handle, bmi *_BITMAPINFO, usage uint32, bits **byte, section syscall.Handle, offset uint32) (bitmap syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall6(procCreateDIBSection.Addr(), 6, uintptr(dc), uintptr(unsafe.Pointer(bmi)), uintptr(usage), uintptr(unsafe.Pointer(bits)), uintptr(section), uintptr(offset))
	bitmap = syscall.Handle(r0)
	if bitmap == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _CreateSolidBrush(color _COLORREF) (brush syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procCreateSolidBrush.Addr(), 1, uintptr(color), 0, 0)
	brush = syscall.Handle(r0)
	if brush == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _DeleteDC(dc syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteDC.Addr(), 1, uintptr(dc), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _DeleteObject(object syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procDeleteObject.Addr(), 1, uintptr(object), 0, 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _FillRect(dc syscall.Handle, rc *_RECT, brush syscall.Handle) (err error) {
	r1, _, e1 := syscall.Syscall(procFillRect.Addr(), 3, uintptr(dc), uintptr(unsafe.Pointer(rc)), uintptr(brush))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _ModifyWorldTransform(dc syscall.Handle, x *_XFORM, mode uint32) (err error) {
	r1, _, e1 := syscall.Syscall(procModifyWorldTransform.Addr(), 3, uintptr(dc), uintptr(unsafe.Pointer(x)), uintptr(mode))
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _SelectObject(dc syscall.Handle, gdiobj syscall.Handle) (newobj syscall.Handle, err error) {
	r0, _, e1 := syscall.Syscall(procSelectObject.Addr(), 2, uintptr(dc), uintptr(gdiobj), 0)
	newobj = syscall.Handle(r0)
	if newobj == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _SetGraphicsMode(dc syscall.Handle, mode int32) (oldmode int32, err error) {
	r0, _, e1 := syscall.Syscall(procSetGraphicsMode.Addr(), 2, uintptr(dc), uintptr(mode), 0)
	oldmode = int32(r0)
	if oldmode == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _SetWorldTransform(dc syscall.Handle, x *_XFORM) (err error) {
	r1, _, e1 := syscall.Syscall(procSetWorldTransform.Addr(), 2, uintptr(dc), uintptr(unsafe.Pointer(x)), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _StretchBlt(dcdest syscall.Handle, xdest int32, ydest int32, wdest int32, hdest int32, dcsrc syscall.Handle, xsrc int32, ysrc int32, wsrc int32, hsrc int32, rop uint32) (err error) {
	r1, _, e1 := syscall.Syscall12(procStretchBlt.Addr(), 11, uintptr(dcdest), uintptr(xdest), uintptr(ydest), uintptr(wdest), uintptr(hdest), uintptr(dcsrc), uintptr(xsrc), uintptr(ysrc), uintptr(wsrc), uintptr(hsrc), uintptr(rop), 0)
	if r1 == 0 {
		if e1 != 0 {
			err = errnoErr(e1)
		} else {
			err = syscall.EINVAL
		}
	}
	return
}

func _GetDeviceCaps(dc syscall.Handle, index int32) (ret int32) {
	r0, _, _ := syscall.Syscall(procGetDeviceCaps.Addr(), 2, uintptr(dc), uintptr(index), 0)
	ret = int32(r0)
	return
}

func _SetProcessDpiAwareness(pdpi uint32) (ret int32) {
	r0, _, _ := syscall.Syscall(procSetProcessDpiAwareness.Addr(), 1, uintptr(pdpi), 0, 0)
	ret = int32(r0)
	return
}

func _GetDpiForWindow(hwnd syscall.Handle) (ret uint32) {
	r0, _, _ := syscall.Syscall(procGetDpiForWindow.Addr(), 1, uintptr(hwnd), 0, 0)
	ret = uint32(r0)
	return
}

func _EnumDisplayDevices(lpdevice uintptr, idevnum uint32, dispdev *_DISPLAY_DEVICE, dwflags uint32) (ok bool) {
	r0, _, _ := syscall.Syscall6(procEnumDisplayDevicesA.Addr(), 4, uintptr(lpdevice), uintptr(idevnum), uintptr(unsafe.Pointer(dispdev)), uintptr(dwflags), 0, 0)
	ok = r0 != 0
	return
}
