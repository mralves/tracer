package tracer

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetStackTrace(t *testing.T) {
	t.Parallel()
	t.Run("when the max depth is specified", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		subject := GetStackTrace(0, 2)
		is.True(len(subject) <= 2, "it should have at most 2 elements")
	})
	t.Run("when the max depth is not specified", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		subject := GetStackTrace(0)
		is.True(len(subject) <= DefaultDepth, "it should have at most 32 elements")
	})
}

func TestStackTrace_String(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	subject := StackTrace{
		{
			Line: 35,
			File: "one.go",
		}, {
			Line: 10,
			File: "two.go",
		}, {
			Line: 15,
			File: "tree.go",
		},
	}
	expected := "one.go:35\n"
	expected += "two.go:10\n"
	expected += "tree.go:15\n"
	is.Equal(expected, subject.String(), "it should return the expected format")
}
