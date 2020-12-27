package notification

import (
	"encoding/json"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/sns"
)

type AWSPublisher struct {
	PlatformApplicationArn string
	SNS                    *sns.SNS
}

type snsMessage struct {
	Default     string `json:"default"`
	APNS        string `json:"APNS"`
	APNSSandbox string `json:"APNS_SANDBOX"`
	GCM         string `json:"GCM"`
	ADM         string `json:"ADM"`
}

type apns struct {
	Aps aps `json:"aps"`
}

type aps struct {
	Alert alert `json:"alert"`
}

type alert struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

type gcm struct {
	Data data `json:"data"`
}
type adm struct {
	Data data `json:"data"`
}

type data struct {
	Message string `json:"message"`
	URL     string `json:"url"`
}

func NewAWS(
	accessKey string,
	secretKey string,
	arn string,
) AWSPublisher {
	sess := session.Must(session.NewSession(&aws.Config{
		Credentials: credentials.NewStaticCredentials(
			accessKey, secretKey, "",
		),
		Region: aws.String("ap-northeast-1"),
	}))
	svc := sns.New(sess)
	return AWSPublisher{
		PlatformApplicationArn: arn,
		SNS: svc,
	}
}

func (a AWSPublisher) Send(input Input) error {
	log.Println(input)
	edarn, err := a.createEndpoint(input.DeviceToken)
	if err != nil {
		return err
	}
	msg, err := a.newMessage(input.Title, input.Message, input.URL)
	if err != nil {
		return err
	}
	log.Println(msg)

	return a.publish(edarn, msg)
}

func (a AWSPublisher) publish(targetArn string, message string) error {
	ipt := &sns.PublishInput{
		MessageStructure: aws.String("json"),
		TargetArn:        aws.String(targetArn),
		Message:          aws.String(message),
	}
	_, err := a.SNS.Publish(ipt)
	return err
}

func (a AWSPublisher) newMessage(title, message, url string) (string, error) {

	apnsdata := apns{
		Aps: aps{
			Alert: alert{
				Title: title,
				Body:  message,
			},
		},
	}

	apnsmsg, err := json.Marshal(apnsdata)
	if err != nil {
		return "", err
	}

	gcmdata := gcm{
		Data: data{
			Message: message,
			URL:     url,
		},
	}
	gcmmsg, err := json.Marshal(gcmdata)
	if err != nil {
		return "", err
	}

	admdata := adm{
		Data: data{
			Message: message,
			URL:     url,
		},
	}
	admmsg, err := json.Marshal(admdata)
	if err != nil {
		return "", err
	}

	snsmsg := snsMessage{
		Default:     message,
		APNS:        string(apnsmsg),
		APNSSandbox: string(apnsmsg),
		GCM:         string(gcmmsg),
		ADM:         string(admmsg),
	}

	msg, err := json.Marshal(snsmsg)
	if err != nil {
		return "", err
	}

	return string(msg), nil
}

func (a AWSPublisher) createEndpoint(deviceToken string) (string, error) {
	ipt := &sns.CreatePlatformEndpointInput{
		PlatformApplicationArn: &a.PlatformApplicationArn,
		Token: &deviceToken,
	}

	output, err := a.SNS.CreatePlatformEndpoint(ipt)
	if err != nil {
		return "", err
	}

	eparn := output.EndpointArn

	return *eparn, err
}
