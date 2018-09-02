package tracer

import (
	"fmt"
	"io"
	"os"
	"time"
)

type Entry struct {
	Level         uint8
	Message       string
	Args          []interface{}
	Time          time.Time
	TransactionId string
	StackTrace    StackTrace
}

type Writer interface {
	Write(entry Entry)
}

type Formatter func(entry Entry) string

type fileWriter struct {
	formatter Formatter
	writer    io.Writer
}

func (fw *fileWriter) Write(entry Entry) {
	_, err := io.WriteString(fw.writer, fw.formatter(entry))
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL ON LOGGING ENTRY BECAUSE: %v\n", err.Error())
	}
}

func NewFileWriter(file io.Writer, formatter Formatter) Writer {
	return &fileWriter{
		writer:    file,
		formatter: formatter,
	}
}

func SimpleFormatter(format string) Formatter {
	return func(entry Entry) string {

	}
}
