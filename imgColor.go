package dmsoft

import ole "github.com/go-ole/go-ole"

func (com *DmSoft) AppendPicAddr(picInfo string, addr, size int) string {
	ret, _ := com.dm.CallMethod("AppendPicAddr", picInfo, addr, size)
	return ret.ToString()
}

func (com *DmSoft) Capture(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod("Capture", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *DmSoft) CaptureGif(x1, y1, x2, y2 int, file string, delay, time int) int {
	ret, _ := com.dm.CallMethod("CaptureGif", x1, y1, x2, y2, file, delay, time)
	return int(ret.Val)
}

func (com *DmSoft) CaptureJpg(x1, y1, x2, y2 int, file string, quality int) int {
	ret, _ := com.dm.CallMethod("CaptureJpg", x1, y1, x2, y2, file, quality)
	return int(ret.Val)
}

func (com *DmSoft) CapturePng(x1, y1, x2, y2 int, file string) int {
	ret, _ := com.dm.CallMethod("CapturePng", x1, y1, x2, y2, file)
	return int(ret.Val)
}

func (com *DmSoft) CapturePre(file string) int {
	ret, _ := com.dm.CallMethod("CapturePre", file)
	return int(ret.Val)
}

func (com *DmSoft) CmpColor(x int, y int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("CmpColor", x, y, color, sim)
	return int(ret.Val)
}

func (com *DmSoft) EnableDisplayDebug(enableDebug int) int {
	ret, _ := com.dm.CallMethod("EnableDisplayDebug", enableDebug)
	return int(ret.Val)
}

func (com *DmSoft) EnableFindPicMultithread(enable int) int {
	ret, _ := com.dm.CallMethod("EnableFindPicMultithread", enable)
	return int(ret.Val)
}

func (com *DmSoft) EnableGetColorByCapture(enable int) int {
	ret, _ := com.dm.CallMethod("EnableGetColorByCapture", enable)
	return int(ret.Val)
}

func (com *DmSoft) FindColor(x1, y1, x2, y2 int, color string, sim float32, dir int, intX *int, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColor", x1, y1, x2, y2, color, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DmSoft) FindColorBlock(x1, y1, x2, y2 int, color string, sim float32, count, width, height int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindColorBlock", x1, y1, x2, y2, color, sim, count, width, height, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DmSoft) FindColorBlockEx(x1, y1, x2, y2 int, color string, sim float32, count, width, height int) string {
	ret, _ := com.dm.CallMethod("FindColorBlockEx", x1, y1, x2, y2, color, sim, count, width, height)
	return ret.ToString()
}

func (com *DmSoft) FindColorE(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorE", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindColorEx(x1, y1, x2, y2 int, color string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindColorEx", x1, y1, x2, y2, color, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindMulColor(x1, y1, x2, y2 int, color string, sim float32) int {
	ret, _ := com.dm.CallMethod("FindMulColor", x1, y1, x2, y2, color, sim)
	return int(ret.Val)
}

func (com *DmSoft) FindMultiColor(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindMultiColor", x1, y1, x2, y2, firstColor, offsetColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DmSoft) FindMultiColorE(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorE", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindMultiColorEx(x1, y1, x2, y2 int, firstColor string, offsetColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindMultiColorEx", x1, y1, x2, y2, firstColor, offsetColor, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindPic(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) int {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPic", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return int(ret.Val)
}

func (com *DmSoft) FindPicE(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicE", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindPicEx(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicEx", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindPicExS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int) string {
	ret, _ := com.dm.CallMethod("FindPicExS", x1, y1, x2, y2, picName, deltaColor, sim, dir)
	return ret.ToString()
}

func (com *DmSoft) FindPicS(x1, y1, x2, y2 int, picName string, deltaColor string, sim float32, dir int, intX, intY *int) string {
	x := ole.NewVariant(ole.VT_I4, 0)
	y := ole.NewVariant(ole.VT_I4, 0)
	ret, _ := com.dm.CallMethod("FindPicS", x1, y1, x2, y2, picName, deltaColor, sim, dir, &x, &y)
	*intX = int(x.Val)
	*intY = int(y.Val)
	x.Clear()
	y.Clear()
	return ret.ToString()
}

func (com *DmSoft) GetAveHSV(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("GetAveHSV", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *DmSoft) GetAveRGB(x1, y1, x2, y2 int) string {
	ret, _ := com.dm.CallMethod("GetAveRGB", x1, y1, x2, y2)
	return ret.ToString()
}

func (com *DmSoft) GetColor(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColor", x, y)
	return ret.ToString()
}

func (com *DmSoft) GetColorHSV(x, y int) string {
	ret, _ := com.dm.CallMethod("GetColorHSV", x, y)
	return ret.ToString()
}

func (com *DmSoft) GetPicSize(picName string) string {
	ret, _ := com.dm.CallMethod("GetPicSize", picName)
	return ret.ToString()
}

func (com *DmSoft) IsDisplayDead(x1, y1, x2, y2, time int) int {
	ret, _ := com.dm.CallMethod("IsDisplayDead", x1, y1, x2, y2, time)
	return int(ret.Val)
}

func (com *DmSoft) MatchPicName(picName string) string {
	ret, _ := com.dm.CallMethod("MatchPicName", picName)
	return ret.ToString()
}

func (com *DmSoft) SetExcludeRegion(mode int, info string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", mode, info)
	return int(ret.Val)
}

func (com *DmSoft) SetPicPwd(pwd string) int {
	ret, _ := com.dm.CallMethod("SetExcludeRegion", pwd)
	return int(ret.Val)
}
