package service

import "study5.1-dip/notifier"

type NotificationService struct {
	sender notifier.Notifier
}

func NewNotificationService(sender notifier.Notifier) *NotificationService {
	return &NotificationService{sender: sender}
}

func (n *NotificationService) Alert(message string) {
	n.sender.Send(message)
}
