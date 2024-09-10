package main

import (
	"runtime"
	"testing"
)

func TestOnlyOnMac(t *testing.T) {
	if runtime.GOOS == "darwin" {
		onlyOnMac()
	} else {
		onlyOnLinux()
	}
}
