package transport

import "github.com/ruraomsk/ag-server/logger"

func senderReplay() {
	for {
		s := <-Sender
		logger.Debug.Printf("%v", s)
	}
}
