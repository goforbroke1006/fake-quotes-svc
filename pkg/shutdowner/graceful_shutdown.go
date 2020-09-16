package shutdowner

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func WaitTerminateSignal(cb func()) {
	done := make(chan bool)

	signals := make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		fmt.Println("press Ctrl+C to terminate...")
		<-signals
		fmt.Println("program termination...")

		cb()

		done <- true
	}()

	<-done
	fmt.Println("bye-bye...")
}
