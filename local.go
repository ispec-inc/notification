package notification

import "github.com/davecgh/go-spew/spew"

// LocalLogger is a logginng the notification information in console.
// Thisis an implementation of the `Service` interface
type LocalLogger struct {
}

func NewLocal() LocalLogger {
	return LocalLogger{}
}

func (l LocalLogger) Send(input Input) error {
	spew.Dump(input)

	return nil
}
