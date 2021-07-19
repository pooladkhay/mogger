package mogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Mogger interface {
	AddSubService(name string) Mogger
	LogToFile(logLevel LogLevel, message string)
	LogToStdErr()
	LogToFileAndStdOut()
}

type mogger struct {
	ServiceName string
	SubService  string
	OutputPath  string
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

func (m mogger) AddSubService(name string) Mogger {
	return &mogger{
		ServiceName: m.ServiceName,
		SubService:  name,
		OutputPath:  m.OutputPath,
	}
}

func (m mogger) LogToFile(logLevel LogLevel, message string) {
	logMessage := fmt.Sprintf("[%s][%s] - %s", time.Now().Format(time.RFC3339), logLevel, message)
	parentFileAddr := fmt.Sprintf("/var/log/%s", m.ServiceName)

	if m.SubService != "" {
		logMessage = fmt.Sprintf("[%s][%s][%s] - %s", m.SubService, time.Now().Format(time.RFC3339), logLevel, message)
	}

	if m.OutputPath != "" {
		parentFileAddr = fmt.Sprintf("%s/%s", m.OutputPath, m.ServiceName)
		if m.SubService != "" {
			parentFileAddr = fmt.Sprintf("%s/%s/%s", m.OutputPath, m.ServiceName, m.SubService)
		}
	}

	if _, err := os.Stat(parentFileAddr); os.IsNotExist(err) {
		err := os.MkdirAll(parentFileAddr, 0700)
		if err != nil {
			log.Fatalln("err creating dirs: ", err)
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/%s-%s.log", parentFileAddr, logLevel, time.Now().Format(time.RFC3339)), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("err opening file: ", err)
	}
	defer f.Close()
	_, err = f.WriteString(logMessage)
	if err != nil {
		log.Fatalln("err writing to file: ", err)
	}
}
func (m mogger) LogToStdErr()        {}
func (m mogger) LogToFileAndStdOut() {}
