package main

import (
	"fmt"
	"time"
	"os2-1/FileWrite"
)

func main() {
	now := time.Now()
	time.Sleep(time.Second*3)
	filewrite.Fwrite()
	fmt.Printf("経過:%vms\n", time.Since(now).Milliseconds())
}
