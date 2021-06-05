package RUNPE

import (
	"fmt"
	"os"
)

func RUNPE() {
	if len(os.Args) < 4 {
		fmt.Println("Please input payload and target path!")
		os.Exit(1)
	}
	payloadPath := os.Args[1]
	targetPath := os.Args[2]
	arguments := os.Args[3]
	HollowProcess(payloadPath, targetPath, arguments)
}
