package tracer

import (
	"runtime"
)

const DefaultDepth = 32

type Caller struct {
	File string
	Line int
}

type StackTrace []Caller

func GetStackTrace(skip int, maxDepth ...int) StackTrace {
	maxDepth = append(maxDepth, DefaultDepth)
	depth := maxDepth[0] + skip
	var stack StackTrace
	for i := skip; i < depth; i++ {
		_, file, line, ok := runtime.Caller(i)
		if !ok {
			break
		}
		stack = append(stack, Caller{
			File: file,
			Line: line,
		})

	}
	return stack
}
