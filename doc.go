// Package notification provides a system by which it is possible to mock your objects
// and verify calls are happening as expected.
//
// Example Usage
// n := notification.NewAWS(
//     "YOUR AWS Access Key",
//     "YOUR AWS Secret Key",
//     "YOUR AWS Platform Application ARN",
// )
// ipt := notification.Input{
//     Title: "Hello",
//     Message: "Message",
//     DeviceToken: "iOS Device Token",
// }
// err := n.Send(ipt)
// if err != nil {
//     //handling error
// }
//
package notification
