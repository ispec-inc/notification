# notification

# Background
ネイティブのアプリ向けに通知を実装する際には、AWSやFirebaseなどのクラウドサービスを使用するケースが多い。
しかし、その通知サービスを扱うには

- 各クラウドサービスの仕様理解に時間がかかってしまう。
- 外部サービスに依存してしまうため、テスタビリティの高い実装が難しい

などの問題がある。

# Objective
高いテスタビリティと、高いユーザービリティ(仕様理解への時間を下げたシンプルなインターフェイス)を兼ね備えた通知のライブラリを構築する。

# High-level Architecture
![notifications](https://user-images.githubusercontent.com/20740529/104127036-317d4500-53a3-11eb-9220-4a57eacb464f.jpg)

## Interface-Driven
`notification.Service`というInterface経由で呼び出しをする設計をしている。これにより、ローカル・テスト環境では`notification.LocalLogger`をDIし、ステージング・本番環境では`notification.AWSPublisher`をDIすることで、高いテスタビリティ、開発体験を維持できる。


# Implementation
## LocalLogger
ローカル環境・テスト環境でDIすることを前提としたStruct。`notificaion.Input`をdumpしてログに出力をする。

## AWSPublisher
[AWS SNS](https://aws.amazon.com/jp/sns/)を用いた通知の実装。内部的にDeviceTokenを投げてEndpointを作成し、そのEndpointに対して通知を送信する実装にした。
Subscriptionなどを用いた実装になっているため、一回に大量の通知を送らなければいけない場合などは、

- bulkに送信できるメソッドを実装する
- Input側で複数のDeviceTokenを受け取れるようにする

などの方法で実装を検討する。

## FCM
TODO
