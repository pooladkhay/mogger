package main

import "github.com/pooladkhay/mogger"

func main() {
	m := mogger.NewMogger("mogger_example", "./logs")

	// Log without subServiceName
	m.LogToFile(mogger.Debug, "without subServiceName")

	// Log with subServiceName
	log := m.AddSubService("Main_func")
	log.LogToFile(mogger.Info, "this is an info log")
	log.LogToFile(mogger.Warn, "this is a warn log")

	anotherFunc(m)
}

func anotherFunc(m mogger.Mogger) {
	// Log with subServiceName
	log := m.AddSubService("Another_func")
	log.LogToFile(mogger.Debug, "this is a debug log")
	log.LogToFile(mogger.Fatal, "this is a fatal log")
}
