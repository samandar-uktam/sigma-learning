package notifier

// This is an Abstract Factory that is creating families of objects such as email, sms and push
type NotificationFactory interface {
	CreateEmailSender() EmailSender
	CreateSmsSender() SmsSender
	CreatePushSender() PushSender
}

type ProdNotificationFactory struct{}

func (p *ProdNotificationFactory) CreateEmailSender() EmailSender {
	return &ProdEmailSender{}
}

func (p *ProdNotificationFactory) CreateSmsSender() SmsSender {
	return &ProdSmsSender{}
}

func (p *ProdNotificationFactory) CreatePushSender() PushSender {
	return &ProdPushSender{}
}

type DevNotificationFactory struct{}

func (p *DevNotificationFactory) CreateEmailSender() EmailSender {
	return &ProdEmailSender{}
}

func (p *DevNotificationFactory) CreateSmsSender() SmsSender {
	return &ProdSmsSender{}
}

func (p *DevNotificationFactory) CreatePushSender() PushSender {
	return &ProdPushSender{}
}
