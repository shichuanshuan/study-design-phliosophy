package main

import (
	"study5.1-dip/notifier"
	"study5.1-dip/service"
)

func main() {
	email := &notifier.EmailNotifier{}
	sms := &notifier.SMSNotifier{}

	// 使用 Email 通知
	emailService := service.NewNotificationService(email)
	emailService.Alert("System down!")

	// 使用 SMS 通知
	smsService := service.NewNotificationService(sms)
	smsService.Alert("CPU usage high!")
}
