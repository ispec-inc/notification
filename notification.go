package notification

// Service is an interface of notification.
type Service interface {
	Send(input Input) error
}

// Input is a struct which input to `Send` method of `Service` interface.
type Input struct {
	Title       string
	Message     string
	URL         string
	DeviceToken string
}
