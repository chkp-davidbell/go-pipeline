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

func TestOnlyOnMac23(t *testing.T) {
	if runtime.GOOS == "darwin" {
		onlyOnMac()
	} else {
		onlyOnLinux()
	}
}
func TestOnlyOnMac3(t *testing.T) {
	if runtime.GOOS == "darwin" {
		onlyOnMac()
	} else {
		onlyOnLinux()
	}
}
func TestOnlyOnMac4(t *testing.T) {
	if runtime.GOOS == "darwin" {
		onlyOnMac()
	} else {
		onlyOnLinux()
	}
}
