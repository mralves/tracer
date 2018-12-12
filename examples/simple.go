package main

import (
	"github.com/mralves/tracer"
	"os"
	"time"
)

type RelevantInformation struct {
	Message string
	Time    time.Time
}

func main() {
	tracer.RegisterWriter(tracer.NewFileWriter(os.Stdout, tracer.SimpleFormatter("@message\n", nil)))
	logger := tracer.GetLogger("simple")
	logger.Info("{{(arg 0).Message}} at {{(arg 0).Time}}", RelevantInformation{Message: "oh no", Time: time.Now()})
}
