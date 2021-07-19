package mogger

import (
	"fmt"
	"log"
	"os"
	"time"
)

type Mogger interface {
	AddSubService(name string) Mogger
	saveToFile(logLevel LogLevel, message string)

	Info(message string)
	Debug(message string)
	Warn(message string)
	Fatal(message string)

	InfoAndToFile(message string)
	DebugAndToFile(message string)
	WarnAndToFile(message string)
	FatalAndToFile(message string)
}

type mogger struct {
	ServiceName string
	SubService  string
	OutputPath  string
}

// Returns a new Mogger instance.
// serviceName is required.
// outputPath is optional (Default: /usr/local/var/log/{serviceName}).
func New(serviceName string, outputPath ...string) Mogger {
	if serviceName == "" {
		log.Fatalln("serviceName is required.")
	}
	if len(outputPath) != 0 {
		return &mogger{
			ServiceName: serviceName,
			OutputPath:  outputPath[0],
		}
	}
	return &mogger{
		ServiceName: serviceName,
		OutputPath:  fmt.Sprintf("/usr/local/var/log/%s", serviceName),
	}
}

// Use sub-service names in your different packages and functions to differentiate the origin of logs.
func (m mogger) AddSubService(name string) Mogger {
	return &mogger{
		ServiceName: m.ServiceName,
		SubService:  name,
		OutputPath:  m.OutputPath,
	}
}

// Prints a info log message to standard error(stderr).
func (m mogger) Info(message string) {
	logMessage := fmt.Sprintf("%s%s%s %s %s%s%s %s %s",
		InfoColor("["), time.Now().Format(time.RFC3339), InfoColor("]"),
		InfoColor("-"),
		InfoColor("["), InfoColor(Info), InfoColor("]"),
		InfoColor("->"),
		message,
	)
	if m.SubService != "" {
		logMessage = fmt.Sprintf("%s%s%s %s %s%s%s %s %s%s%s %s %s",
			InfoColor("["), m.SubService, InfoColor("]"),
			InfoColor("-"),
			InfoColor("["), time.Now().Format(time.RFC3339), InfoColor("]"),
			InfoColor("-"),
			InfoColor("["), InfoColor(Info), InfoColor("]"),
			InfoColor("->"),
			message,
		)
	}
	fmt.Fprintln(os.Stderr, logMessage)
}

// Prints a debug log message to standard error(stderr).
func (m mogger) Debug(message string) {
	logMessage := fmt.Sprintf("%s%s%s %s %s%s%s %s %s",
		DebugColor("["), time.Now().Format(time.RFC3339), DebugColor("]"),
		DebugColor("-"),
		DebugColor("["), DebugColor(Debug), DebugColor("]"),
		DebugColor("->"),
		message,
	)
	if m.SubService != "" {
		logMessage = fmt.Sprintf("%s%s%s %s %s%s%s %s %s%s%s %s %s",
			DebugColor("["), m.SubService, DebugColor("]"),
			DebugColor("-"),
			DebugColor("["), time.Now().Format(time.RFC3339), DebugColor("]"),
			DebugColor("-"),
			DebugColor("["), DebugColor(Debug), DebugColor("]"),
			DebugColor("->"),
			message,
		)
	}
	fmt.Fprintln(os.Stderr, logMessage)
}

// Prints a warn log message to standard error(stderr).
func (m mogger) Warn(message string) {
	logMessage := fmt.Sprintf("%s%s%s %s %s%s%s %s %s",
		WarnColor("["), time.Now().Format(time.RFC3339), WarnColor("]"),
		WarnColor("-"),
		WarnColor("["), WarnColor(Warn), WarnColor("]"),
		WarnColor("->"),
		message,
	)
	if m.SubService != "" {
		logMessage = fmt.Sprintf("%s%s%s %s %s%s%s %s %s%s%s %s %s",
			WarnColor("["), m.SubService, WarnColor("]"),
			WarnColor("-"),
			WarnColor("["), time.Now().Format(time.RFC3339), WarnColor("]"),
			WarnColor("-"),
			WarnColor("["), WarnColor(Warn), WarnColor("]"),
			WarnColor("->"),
			message,
		)
	}
	fmt.Fprintln(os.Stderr, logMessage)
}

// Prints a fatal log message to standard error(stderr) and exits the program with non-zero status code.
func (m mogger) Fatal(message string) {
	logMessage := fmt.Sprintf("%s%s%s %s %s%s%s %s %s",
		FatalColor("["), time.Now().Format(time.RFC3339), FatalColor("]"),
		FatalColor("-"),
		FatalColor("["), FatalColor(Fatal), FatalColor("]"),
		FatalColor("->"),
		message,
	)
	if m.SubService != "" {
		logMessage = fmt.Sprintf("%s%s%s %s %s%s%s %s %s%s%s %s %s",
			FatalColor("["), m.SubService, FatalColor("]"),
			FatalColor("-"),
			FatalColor("["), time.Now().Format(time.RFC3339), FatalColor("]"),
			FatalColor("-"),
			FatalColor("["), FatalColor(Fatal), FatalColor("]"),
			WarnColor("->"),
			message,
		)
	}
	fmt.Fprintln(os.Stderr, logMessage)
	os.Exit(1)
}

// Prints a info log message to standard error(stderr) and saves it to a file.
func (m mogger) InfoAndToFile(message string) {
	m.Info(message)
	m.saveToFile(Info, message)
}

// Prints a debug log message to standard error(stderr) and saves it to a file.
func (m mogger) DebugAndToFile(message string) {
	m.Debug(message)
	m.saveToFile(Debug, message)
}

// Prints a warn log message to standard error(stderr) and saves it to a file.
func (m mogger) WarnAndToFile(message string) {
	m.Warn(message)
	m.saveToFile(Warn, message)
}

// Prints a fatal log message to standard error(stderr), Saves it to a file and exits the program with non-zero status code.
func (m mogger) FatalAndToFile(message string) {
	m.saveToFile(Fatal, message)
	m.Fatal(message)
}

func (m mogger) saveToFile(logLevel LogLevel, message string) {
	logMessage := fmt.Sprintf("[%s] - [%s] -> %s\n", time.Now().Format(time.RFC3339), logLevel, message)
	outputPath := m.OutputPath

	if m.SubService != "" {
		logMessage = fmt.Sprintf("[%s] - [%s] - [%s] -> %s\n", m.SubService, time.Now().Format(time.RFC3339), logLevel, message)
		outputPath = fmt.Sprintf("%s/%s", m.OutputPath, m.SubService)
	}

	if _, err := os.Stat(outputPath); os.IsNotExist(err) {
		err := os.MkdirAll(outputPath, 0700)
		if err != nil {
			log.Fatalln("err creating dirs: ", err)
		}
	}

	f, err := os.OpenFile(fmt.Sprintf("%s/%s-%s.log", outputPath, logLevel, time.Now().Format(time.RFC3339)), os.O_CREATE|os.O_WRONLY, os.ModePerm)
	if err != nil {
		log.Fatalln("err opening file: ", err)
	}
	defer f.Close()
	_, err = f.WriteString(logMessage)
	if err != nil {
		log.Fatalln("err writing to file: ", err)
	}
}
