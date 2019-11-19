package tracer

import (
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewContext(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	subject := NewContext(Debug)
	is.NotNil(subject, "it should not be nil")
}

func TestContext_ChildContext(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	subject := NewContext(Warning).(*context)
	child := subject.ChildContext("owner").(*context)
	is.NotNil(child, "it should not be nil")
	is.Contains(subject.children, "owner", "it should update it's hash o children")
	is.Equal(subject.minimumLevel, child.minimumLevel, "it should have the same minimumLevel as it's parent")
}

func TestContext_GetLogger(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	subject := NewContext(Warning).(*context)
	loggerA := subject.GetLogger("A")
	loggerB := subject.GetLogger("B")
	loggerB2 := subject.GetLogger("B")
	is.NotEqual(loggerA, loggerB, "it should return two different loggers")
	is.Equal(loggerB, loggerB2, "it should return the same logger")
}

type writerM struct {
	mock.Mock
}

func (wm writerM) Write(entry Entry) {
	wm.Called(entry)
}

func TestContext_RegisterWriter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	writer := writerM{}
	subject := NewContext(Warning).(*context)
	subject.RegisterWriter(writer)
	is.Len(subject.writers, 1)
}

func TestContext_OverwriteChildren(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	subject := NewContext(Warning).(*context)
	child := subject.ChildContext("owner")
	child.MinimumLevel(Fatal)
	is.Equal(Fatal, child.GetMinimumLevel(), "it should set the minimalLevel")
	subject.MinimumLevel(Debug)
	is.Equal(Debug, subject.GetMinimumLevel(), "it should set the minimalLevel")
	subject.OverwriteChildren()
	is.Equal(subject.GetMinimumLevel(), child.GetMinimumLevel(), "it should change the minimalLevel")
}
