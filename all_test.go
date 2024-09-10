package main

import (
	"runtime"
	"testing"
)

func TestOnlyOnMac2(t *testing.T) {
	if runtime.GOOS == "darwin" {
		onlyOnMac()
	} else {
		onlyOnLinux()
	}
}

func TestOnlyOnLinux(t *testing.T) {
	if runtime.GOOS == "linux" {
		onlyOnLinux()
	} else {
		onlyOnMac()
	}
}
