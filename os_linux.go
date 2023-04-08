//go:build linux

// Package log +build linux
package log

func init() {
	ResetColor = "\033[0m"
	Red = "\033[31m"
	Green = "\033[32m"
	Yellow = "\033[33m"
}
