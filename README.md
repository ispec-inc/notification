notification
---
[![Go Report Card](https://goreportcard.com/badge/github.com/ispec-inc/notification)](https://goreportcard.com/report/github.com/ispec-inc/notification) [![PkgGoDev](https://pkg.go.dev/badge/github.com/ispec-inc/notification)](https://pkg.go.dev/github.com/ispec-inc/notification)

A simple notification library written in golang.

## Installation
```bash
$ go get -u github.com/ispec-inc/notification
```

## Usage

```
n := notification.NewAWS(
    "YOUR AWS Access Key",
    "YOUR AWS Secret Key",
    "YOUR AWS Platform Application ARN",
)
ipt := notification.Input{
    Title: "Hello",
    Message: "Message",
    DeviceToken: "iOS Device Token",
}
err := n.Send(ipt)
if err != nil {
    //handling error
}
```

In the test time, you can use `mock_notification` package.
Also in the development time, you can use `LocalPublisher` which logging the notification information in console.

## LICENCE
Copyright © 2017 Yusuke Yamada MIT license
