package main

import "runtime"

func main() {
	//check if the os is mac
	if runtime.GOOS == "darwin" {
		onlyOnMac()
	} else {
		onlyOnLinux()
	}
}
