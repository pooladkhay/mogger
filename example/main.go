package main

import (
	"github.com/pooladkhay/mogger"
)

func main() {
	m := mogger.New("mogger_example", "./logs")

	// log without sub-service name
	m.InfoAndToFile("from main without sub-service name")

	// log with a sub-service name
	log := m.AddSubService("main_func")
	log.Info("info log from main func")
	log.Debug("debug log from main func")

	// -AndToFile prints to stderr and saves it to a file.
	log.WarnAndToFile("warn log from main func")

	anotherFunc(m)
}

func anotherFunc(m mogger.Mogger) {
	log := m.AddSubService("another_func")
	log.DebugAndToFile("debug log from another func")
	// fatal causes program to exit with a non-zero status code.
	log.FatalAndToFile("fatal log from main func")
}
