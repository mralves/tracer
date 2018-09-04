# Tracer
Easy to use and extend log library

## How to install
Using go get (not recommended):
```bash
go get github.com/mralves/tracer
```

Using [dep](github.com/golang/dep) (recommended):
```bash
dep ensure --add github.com/stretchr/testify@<version>
```

## How to use

Below follows a simple example of how to use this lib:

```go
package main

import "github.com/mralves/tracer"

func main() {
    logger := tracer.GetLogger("moduleA")
    logger.AutoTrace(true) // tells the logger to automatic creates transactions using the first optional argument on a log
    logger.Info("transaction A", "A")
    logger.Info("transaction B", "B")
    logger.Info("transaction B", "B")
    logger.Info("transaction A", "A")
    logger.Info("transaction A", "A")
    logger = logger.Trace("C") // now all logs on this logger will be on the transaction C
    logger.Info("transaction C")
    logger.Info("transaction C", "A")
}
```

