package notifier

func getNotificationFactory(env string) NotificationFactory {
	if env == "prod" {
		return &ProdNotificationFactory{}
	}
	return &DevNotificationFactory{}
}

func main() {
	env := "dev"
	factory := getNotificationFactory(env)

	emailSender := factory.CreateEmailSender()
	smsSender := factory.CreateSmsSender()
	pushSender := factory.CreatePushSender()

	emailSender.Send("example@gmail.com", "Hello email!")
	smsSender.Send("+1234567890", "Hello sms!")
	pushSender.Send("device-123", "Hello push!")
}
