package main

import "github.com/pooladkhay/mogger"

func main() {
	m := mogger.NewMogger("my_service_name", "./logs")

	m.LogToFile(mogger.Info, "this is an info log")
	m.LogToFile(mogger.Debug, "this is a debug log")
	m.LogToFile(mogger.Warn, "this is a warn log")
	m.LogToFile(mogger.Fatal, "this is a fatal log")
}
