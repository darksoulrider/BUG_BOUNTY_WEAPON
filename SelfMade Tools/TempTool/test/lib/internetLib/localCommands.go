package internetlib

import (
	"fmt"
	"os/exec"
)

func RunLocalCmd(cmd string, st string, dir string) {
	fmt.Println("Running a Command...")

	myCmd := exec.Command(cmd, st, dir)
	output, err := myCmd.Output()
	if err != nil {
		fmt.Println("Error occurred")
		fmt.Println(err.Error())
		return
	}

	fmt.Println(string(output))

}
