package log

import (
	"fmt"
	"io"
	"os"
)

type Level int

const (
	LevelInfo Level = iota
	LevelWarn
	LevelError
	LevelFatal
)

var (
	CurrentLevel Level     = LevelInfo
	CurrentOut   io.Writer = os.Stderr
)

func Print(level Level, msg string, args ...any) {
	if level < CurrentLevel {
		return
	}

	fmt.Fprintf(CurrentOut, msg, args...)
	if n := len(msg); n > 0 && msg[n-1] != '\n' {
		fmt.Println()
	}
}

func Info(msg string, args ...any) {
	Print(LevelInfo, msg, args...)
}

func Warn(msg string, args ...any) {
	Print(LevelWarn, msg, args...)
}

func Error(msg string, args ...any) {
	Print(LevelError, msg, args...)
}

func Fatal(msg string, args ...any) {
	Print(LevelFatal, msg, args...)
	os.Exit(1)
}
