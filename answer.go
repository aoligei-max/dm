package dmsoft

func (com *DmSoft) FaqCancel() int {
	ret, _ := com.dm.CallMethod("FaqCancel")
	return int(ret.Val)
}

func (com *DmSoft) FaqCapture(x1, y1, x2, y2, quality, delay, time int) int {
	ret, _ := com.dm.CallMethod("FaqCapture", x1, y1, x2, y2, quality, delay, time)
	return int(ret.Val)
}

func (com *DmSoft) FaqCaptureFromFile(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod("FaqCaptureFromFile", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *DmSoft) FaqCaptureString(str string) int {
	ret, _ := com.dm.CallMethod("FaqCaptureString", str)
	return int(ret.Val)
}

func (com *DmSoft) FaqFetch() string {
	ret, _ := com.dm.CallMethod("FaqFetch")
	return ret.ToString()
}

func (com *DmSoft) FaqGetSize(handle int) int {
	ret, _ := com.dm.CallMethod("FaqGetSize", handle)
	return int(ret.Val)
}

func (com *DmSoft) FaqIsPosted() int {
	ret, _ := com.dm.CallMethod("FaqIsPosted")
	return int(ret.Val)
}

func (com *DmSoft) FaqPost(server string, handle, request_type, time_out int) int {
	ret, _ := com.dm.CallMethod("FaqPost", server, handle, request_type, time_out)
	return int(ret.Val)
}

func (com *DmSoft) FaqSend(server string, handle, request_type, time_out int) string {
	ret, _ := com.dm.CallMethod("FaqSend", server, handle, request_type, time_out)
	return ret.ToString()
}
