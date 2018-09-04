package tracer

import (
	"fmt"
	"sync"
	"time"
)

const (
	Fatal         = uint8(0)
	Alert         = uint8(1)
	Critical      = uint8(2)
	Error         = uint8(3)
	Warning       = uint8(4)
	Notice        = uint8(5)
	Informational = uint8(6)
	Debug         = uint8(7)
)

var defaultWriters []Writer
var lock sync.Locker = &sync.RWMutex{}

type Logger interface {
	Debug(message string, args ...interface{})
	D(message string, args ...interface{})
	Info(message string, args ...interface{})
	I(message string, args ...interface{})
	Warn(message string, args ...interface{})
	W(message string, args ...interface{})
	Error(message string, args ...interface{})
	E(message string, args ...interface{})
	Trace(transactionId string) Logger
	AutoTrace(on bool) Logger
}

type logger struct {
	writers                    []Writer
	createImplicitTransactions bool
	transactionId              string
	owner                      string
}

func GetLogger(owner string) Logger {
	return &logger{
		writers:                    []Writer{},
		owner:                      owner,
		createImplicitTransactions: false,
	}
}

func (l *logger) Debug(message string, args ...interface{}) {
	l.log(Debug, message, args)
}

func (l *logger) D(message string, args ...interface{}) {
	l.log(Debug, message, args)
}

func (l *logger) Info(message string, args ...interface{}) {
	l.log(Informational, message, args)
}

func (l *logger) I(message string, args ...interface{}) {
	l.log(Informational, message, args)
}

func (l *logger) Notice(message string, args ...interface{}) {
	l.log(Notice, message, args)
}

func (l *logger) N(message string, args ...interface{}) {
	l.log(Notice, message, args)
}

func (l *logger) Warn(message string, args ...interface{}) {
	l.log(Warning, message, args)
}

func (l *logger) W(message string, args ...interface{}) {
	l.log(Warning, message, args)
}

func (l *logger) Error(message string, args ...interface{}) {
	l.log(Error, message, args)
}

func (l *logger) E(message string, args ...interface{}) {
	l.log(Error, message, args)
}

func (l *logger) Critical(message string, args ...interface{}) {
	l.log(Critical, message, args)
}

func (l *logger) C(message string, args ...interface{}) {
	l.log(Critical, message, args)
}

func (l *logger) Alert(message string, args ...interface{}) {
	l.log(Alert, message, args)
}

func (l *logger) A(message string, args ...interface{}) {
	l.log(Alert, message, args)
}

func (l *logger) Fatal(message string, args ...interface{}) {
	l.log(Fatal, message, args)
}

func (l *logger) F(message string, args ...interface{}) {
	l.log(Fatal, message, args)
}

func (l *logger) Trace(transactionId string) Logger {
	return &logger{
		writers:                    l.writers,
		createImplicitTransactions: false,
		transactionId:              transactionId,
	}
}

func (l *logger) AutoTrace(on bool) Logger {
	l.createImplicitTransactions = on
	return l
}

func (l *logger) log(level uint8, message string, args []interface{}) {
	transactionId := l.transactionId
	if l.createImplicitTransactions {
		if len(args) > 0 {
			transactionId = fmt.Sprint(args[0])
		}
	}
	entry := Entry{
		Owner:         l.owner,
		Level:         level,
		Message:       message,
		Args:          args,
		Time:          time.Now(),
		TransactionId: transactionId,
		StackTrace:    GetStackTrace(3),
	}
	var wg sync.WaitGroup
	for _, writer := range l.writers {
		wg.Add(1)
		go func(writer Writer, entry Entry, wg *sync.WaitGroup) {
			defer wg.Done()
			writer.Write(entry)
		}(writer, entry, &wg)
	}
	wg.Wait()
}
