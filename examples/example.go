package main

import (
	"log"

	"github.com/davecgh/go-spew/spew"
	"github.com/make-42/go-notifications"
)

func main() {
	notificationReceiver, err := notifications.NewNotificationReceiver(true)
	if err != nil {
		log.Fatal(err)
	}
	notification := notificationReceiver.GetBlocking()
	if notification.Error != nil {
		log.Fatal(err)
	}
	spew.Dump(notification.NotificationBody)
	channel := notificationReceiver.GetChannel()
	for v := range channel {
		if v.Error != nil {
			log.Fatal(v.Error)
		}
		spew.Dump(v.NotificationBody)
	}
}
