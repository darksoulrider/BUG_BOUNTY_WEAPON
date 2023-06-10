package internetlib

import (
	"fmt"
	"os"

	"strings"
)

func Splitfile() {
	filename := "wordlist.txt"

	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file")
		os.Exit(1)
	}
	mck := "hackerone"
	content := strings.Split(string(data), "\n")
	for _, line := range content {
		data := strings.TrimSpace(line)
		fmt.Printf("%s.%s\n", data, mck)
	}
}
