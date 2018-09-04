package tracer

import (
	"bytes"
	"fmt"
	"html/template"
	"io"
	"os"
	"time"
)

type Entry struct {
	Owner         string
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

type FileWriter struct {
	formatter Formatter
	writer    io.Writer
}

func (fw *FileWriter) Write(entry Entry) {
	_, err := fw.writer.Write([]byte(fw.formatter(entry)))
	if err != nil {
		fmt.Fprintf(os.Stderr, "FAIL ON LOGGING ENTRY BECAUSE: %v\n", err.Error())
	}
}

func NewFileWriter(file io.Writer, formatter Formatter) *FileWriter {
	return &FileWriter{
		writer:    file,
		formatter: formatter,
	}
}

func SimpleFormatter(format string) Formatter {
	t, err := template.New("log").Parse(format)
	if err != nil {
		panic(err)
	}
	return func(entry Entry) string {
		var buf bytes.Buffer
		err := t.Execute(&buf, entry)
		if err != nil {
			return fmt.Sprintf("FAILED TO FORMAT ENTRY BECAUSE: %v\n", err.Error())
		}
		return buf.String()
	}
}
