# mogger
Minimalistic Logging Library for Golang

## Installation
```go get github.com/pooladkhay/mogger```

## Colorful Console Output
![mogger](console-output.png?raw=true)

## Usage
```
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
	log.InfoAndToFile("info log from main func")
	log.DebugAndToFile("debug log from main func")
	log.WarnAndToFile("warn log from main func")

	anotherFunc(m)
}

func anotherFunc(m mogger.Mogger) {
	log := m.AddSubService("another_func")
	log.DebugAndToFile("debug log from another func")
	// fatal causes program to exit with a non-zero status code.
	log.FatalAndToFile("fatal log from main func")
}

```

## Contributing
Pull requests are welcomed. For major changes, please open an issue first to discuss what you would like to change.

## License
[MIT](https://github.com/pooladkhay/mogger/blob/main/LICENSE)