package interrupt

import (
	"os"
	"os/signal"
	"syscall"
)

type Interrupt func()

var ShutdownOnSigTerm = func() {
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)
	<-termChan
}
