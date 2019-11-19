package tracer

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewFileWriter(t *testing.T) {
	t.Parallel()
	is := assert.New(t)
	is.NotNil(NewFileWriter(nil, nil), "it should return a valid FileWriter")
}

type mockWriter struct {
	mock.Mock
}

func (mw mockWriter) Write(p []byte) (n int, err error) {
	args := mw.Called(p)
	if args[1] != nil {
		return 0, args[1].(error)
	}
	return args[0].(int), nil
}

func TestFileWriter_Write(t *testing.T) {
	t.Parallel()
	t.Run("when is not possible to write the log", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		called := 0
		expected := Entry{
			Message: "this is a test",
		}
		mw := mockWriter{}
		mw.On("Write", []byte("dummy")).Return(0, errors.New("some-error")).Once()
		subject := &FileWriter{
			writer: mw,
			formatter: func(entry Entry) string {
				called++
				is.Equal(expected, entry, "it should be the expected entry")
				return "dummy"
			},
		}
		is.NotPanics(func() {
			subject.Write(expected)
		}, "it should not panics")
		is.Equal(1, called, "it should be called one time")
		mw.AssertExpectations(t)

	})
	t.Run("when is possible to write the log", func(t *testing.T) {
		t.Parallel()
		is := assert.New(t)
		called := 0
		expected := Entry{
			Message: "this is a test",
		}
		mw := mockWriter{}
		mw.On("Write", []byte("dummy")).Return(15, nil).Once()
		subject := &FileWriter{
			writer: mw,
			formatter: func(entry Entry) string {
				called++
				is.Equal(expected, entry, "it should be the expected entry")
				return "dummy"
			},
		}
		is.NotPanics(func() {
			subject.Write(expected)
		}, "it should not panics")
		is.Equal(1, called, "it should be called one time")
		mw.AssertExpectations(t)
	})
}
