package RUNPE

import (
	"fmt"
	"os"
)

func RUNPE() {
	if len(os.Args) < 3 {
		fmt.Println("Please input payload and target path!")
		os.Exit(1)
	}
	payloadPath := os.Args[1]
	targetPath := os.Args[2]
	HollowProcess(payloadPath, targetPath)
}
