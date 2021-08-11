package logging

import (
	"log"
	"os"
)

var (
	InfoLogger  *log.Logger
	ErrorLogger *log.Logger
)

// Init creates 2 loggers for info and error and makes them globally accessible
func Init() {
	InfoLogger = log.New(os.Stdout, "info: ", 0)
	ErrorLogger = log.New(os.Stderr, "error: ", 0)
}
