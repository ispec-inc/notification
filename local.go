package notification

import (
	"fmt"
	"log"
)

// LocalLogger is a logginng the notification information in console.
// Thisis an implementation of the `Service` interface
type LocalLogger struct {
}

func NewLocal() LocalLogger {
	return LocalLogger{}
}

func (l LocalLogger) Send(input Input) error {
	log.Println("======== notification ========")
	log.Println(fmt.Sprintf("To: %s", input.DeviceToken))
	log.Println(fmt.Sprintf("Title: %s", input.Title))
	log.Println(fmt.Sprintf("Message: %s", input.Message))

	return nil
}
