// Package log provides a global logger intended to be a drop-in for the stdlib log.
package log

import (
	"fmt"
	"os"

	"github.com/kpurdon/zlg"
	"github.com/rs/zerolog"
)

// Logger is the global logger.
var Logger = zlg.New()

func Fatal(v ...interface{}) {
	Logger.Panic(fmt.Sprint(v...))
	os.Exit(1)
}

func Fatalf(format string, v ...interface{}) {
	Logger.Panic(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func Fatalln(v ...interface{}) {
	Logger.Panic(fmt.Sprintln(v...))
	os.Exit(1)
}

func Panic(v ...interface{}) {
	Logger.Panic(fmt.Sprint(v...))
}

func Panicf(format string, v ...interface{}) {
	Logger.Panic(fmt.Sprintf(format, v...))
}

func Panicln(v ...interface{}) {
	Logger.Panic(fmt.Sprintln(v...))
}

func Print(v ...interface{}) {
	Logger.Info(fmt.Sprint(v...))
}

func Printf(format string, v ...interface{}) {
	Logger.Info(fmt.Sprintf(format, v...))
}

func Println(v ...interface{}) {
	Logger.Info(fmt.Sprintln(v...))
}

func With(k string, v interface{}) {
	Logger = Logger.With(k, v)
}

func WithError(err error) {
	Logger = Logger.WithError(err)
}

func Debug(msg string) {
	Logger.Debug(msg)
}

func Info(msg string) {
	Logger.Info(msg)
}

func Error(err error) {
	Logger.Error(err)
}

func WithLevel(lvl zerolog.Level) {
	Logger = Logger.WithLevel(lvl)
}
