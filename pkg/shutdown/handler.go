package shutdown

import (
	"fmt"
	"os"
	"syscall"
)

var GracefulShutdownSignals = []os.Signal{syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL}

func SignalHandle(signal <-chan os.Signal, exit chan<- int, stop func() error, signals ...os.Signal) {
	if len(signals) == 0 {
		signals = GracefulShutdownSignals
	}

	for {
		s := <-signal
		for _, sgnl := range signals {
			if s == sgnl {
				err := stop()
				if err != nil {
					exit <- 1
				} else {
					exit <- 0
				}
				return
			}
		}

		fmt.Printf("Unsupported signal: %s\n", s)
		exit <- -1
	}
}
