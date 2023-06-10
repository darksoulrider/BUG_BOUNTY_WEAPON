package main

import (
	"fmt"
	internetlib "github.com/testCode/lib/internetLib"
)

func main() {
	fmt.Println("Starting....")
	// internetlib.Check_NsLookup("Hackerone.com")
	// internetlib.Splitfile()
	// internetlib.GetWebSourceCode("https://www.facebook.com")
	internetlib.RunLocalCmd("cmd", "/c", "dir")
}
