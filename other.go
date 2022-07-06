package dmsoft

func (com *DmSoft) DmGuard(enable int, types string) int {
	ret, _ := com.dm.CallMethod("DmGuard", enable, types)
	return int(ret.Val)
}

func (com *DmSoft) DmGuardExtract(types, path string) int {
	ret, _ := com.dm.CallMethod("DmGuardExtract", types, path)
	return int(ret.Val)
}

func (com *DmSoft) DmGuardLoadCustom(types, path string) int {
	ret, _ := com.dm.CallMethod("DmGuardLoadCustom", types, path)
	return int(ret.Val)
}

func (com *DmSoft) DmGuardParams(cmd, subcmd, param string) string {
	ret, _ := com.dm.CallMethod("DmGuardParams", cmd, subcmd, param)
	return ret.ToString()
}

func (com *DmSoft) UnLoadDriver() int {
	ret, _ := com.dm.CallMethod("UnLoadDriver")
	return int(ret.Val)
}
