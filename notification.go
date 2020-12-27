package notification

type Service interface {
	Send(input Input) error
}

type Input struct {
	Title       string
	Message     string
	URL         string
	DeviceToken string
}
