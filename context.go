package tracer

import (
	"sync"
)

var DefaultContext Context = nil

func init() {
	DefaultContext = NewContext(Debug, false)
}

type context struct {
	sync.Locker
	parent        Context
	writers       []Writer
	loggers       map[string]Logger
	children      map[string]context
	minimumLevel  uint8
	implicitTrace bool
}

type Context interface {
	GetLogger(owner string) Logger
	RegisterWriter(writer Writer)
	MinimumLevel(level uint8)
	GetMinimumLevel() uint8
	GetWriters() []Writer
	ImplicitTrace(on bool)
	GetImplicitTrace() bool
}

func NewContext(minimumLevel uint8, implicitTransactions bool) Context {
	return &context{
		Locker:        &sync.RWMutex{},
		minimumLevel:  minimumLevel,
		implicitTrace: implicitTransactions,
		writers:       []Writer{},
		loggers:       map[string]Logger{},
		children:      map[string]context{},
		parent:        DefaultContext,
	}
}

func (c *context) newChild() *context {
	child := NewContext(c.minimumLevel, c.implicitTrace).(*context)
	child.parent = c
	return child
}

func (c *context) GetLogger(owner string) Logger {
	c.Lock()
	defer c.Unlock()
	if _, ok := c.loggers[owner]; !ok {
		c.loggers[owner] = &logger{
			Locker:  &sync.RWMutex{},
			Context: c.newChild(),
			owner:   owner,
		}
	}
	return c.loggers[owner]
}

func (c *context) RegisterWriter(writer Writer) {
	c.Lock()
	defer c.Unlock()
	c.writers = append(c.writers, writer)
}

func (c *context) GetWriters() []Writer {
	c.Lock()
	defer c.Unlock()
	var writers []Writer
	if c.parent != nil {
		writers = append(writers, c.parent.GetWriters()...)
	}
	writers = append(writers, c.writers...)
	return writers
}

func (c *context) ImplicitTrace(state bool) {
	c.implicitTrace = state
}

func (c *context) GetImplicitTrace() bool {
	return c.implicitTrace
}

func (c *context) MinimumLevel(level uint8) {
	c.minimumLevel = level
}

func (c *context) GetMinimumLevel() uint8 {
	return c.minimumLevel
}