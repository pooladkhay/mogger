package mogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Mogger interface {
	LogToFile(logLevel LogLevel, message string)
	LogToStdErr()
	LogToFileAndStdOut()
}

type mogger struct {
	// Required.
	ServiceName string
	// Optional. Default: /var/log
	OutputPath string
}

// serviceName is required.
// outputPath is optional (Default: /var/log/{serviceName}).
func NewMogger(serviceName, outputPath string) Mogger {
	if serviceName == "" {
		log.Fatalln("serviceName must be provided.")
	}
	return &mogger{
		ServiceName: serviceName,
		OutputPath:  outputPath,
	}
}

func (m *mogger) LogToFile(logLevel LogLevel, message string) {
	logMessage := fmt.Sprintf("[%s] - %s", logLevel, message)

	parentFileAddr := fmt.Sprintf("/var/log/%s", m.ServiceName)
	if m.OutputPath != "" {
		parentFileAddr = m.OutputPath
	}
	subFileAddr := fmt.Sprintf("%s/%s", parentFileAddr, logLevel)

	if _, err := os.Stat(subFileAddr); os.IsNotExist(err) {
		err := os.MkdirAll(subFileAddr, 0700)
		if err != nil {
			log.Fatalln("err creating dirs: ", err)
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/%s.log", subFileAddr, time.Now().Format(time.RFC3339)), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("err opening file: ", err)
	}
	defer f.Close()
	_, err = f.WriteString(logMessage)
	if err != nil {
		log.Fatalln("err writing to file: ", err)
	}
}
func (m *mogger) LogToStdErr()        {}
func (m *mogger) LogToFileAndStdOut() {}
