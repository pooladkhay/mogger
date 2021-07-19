package main

import (
	"github.com/pooladkhay/mogger"
)

func main() {
	m := mogger.New("mogger_example", "./logs")

	// Log without subServiceName
	m.InfoAndToFile("from main without subServiceName")

	// Log with subServiceName
	log := m.AddSubService("main_func")
	log.InfoAndToFile("info log from main func")
	log.DebugAndToFile("debug log from main func")
	log.WarnAndToFile("warn log from main func")

	anotherFunc(m)
}

func anotherFunc(m mogger.Mogger) {
	// Log with subServiceName
	log := m.AddSubService("another_func")
	log.DebugAndToFile("debug log from another func")

	// fatal causes program to exit with a non-zero status code.
	log.FatalAndToFile("fatal log from main func")
}
