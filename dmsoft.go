// +build windows
// export GOARCH=386

package dmsoft

import (
	"syscall"
	"unsafe"

	ole "github.com/go-ole/go-ole"
	"github.com/go-ole/go-ole/oleutil"
)

var (
	dmReg32         = syscall.NewLazyDLL("DmReg.dll")
	procSetDllPathA = dmReg32.NewProc("SetDllPathA")
	procSetDllPathW = dmReg32.NewProc("SetDllPathW")
)

type DmSoft struct {
	dm       *ole.IDispatch
	IUnknown *ole.IUnknown
}

func New() (dm *DmSoft, err error) {
	var com DmSoft
	ole.CoInitialize(0)
	com.IUnknown, err = oleutil.CreateObject("dm.dmsoft")
	if err != nil {
		return nil, err
	}
	com.dm, err = com.IUnknown.QueryInterface(ole.IID_IDispatch)
	if err != nil {
		return nil, err
	}
	return &com, nil
}

func (com *DmSoft) Release() {
	com.IUnknown.Release()
	com.dm.Release()
	ole.CoUninitialize()
}

func SetDllPathA(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := syscall.Syscall(procSetDllPathA.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	return ret != 0
}

func SetDllPathW(path string, mode int) bool {
	var _p0 *uint16
	_p0, _ = syscall.UTF16PtrFromString(path)
	ret, _, _ := syscall.Syscall(procSetDllPathW.Addr(), 2, uintptr(unsafe.Pointer(_p0)), uintptr(mode), 0)
	return ret != 0
}
