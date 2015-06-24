package mci
import(
	"syscall"
	"unsafe"
)
var winmm = syscall.MustLoadDLL("winmm.dll")
var pmciSendString = winmm.MustFindProc("mciSendStringW")
func mciSendString(cmd string) int {
	r1, _, _ := pmciSendString.Call(uintptr(unsafe.Pointer(syscall.StringToUTF16Ptr(cmd))), 0, 0, 0)
	return int(r1)
}
func Send(cmd string) {
	mciSendString(cmd)
}
