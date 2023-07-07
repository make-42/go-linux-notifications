package notifications

import (
	"errors"

	"github.com/godbus/dbus/v5"
)

const (
	notificationDBusObjectPath = "/org/freedesktop/Notifications"
	notificationDBusInterface  = "org.freedesktop.Notifications"
	notificationDBusMember     = "Notify"
)

// This function creates a new `NotificationReceiver` object that can then be used to receive notifications. This object has to be closed with the `Close()` method.
func NewNotificationReceiver(systemRepeatsTwice bool) (NotificationReceiver, error) {
	conn, err := dbus.ConnectSessionBus()
	if err != nil {
		return NotificationReceiver{}, err
	}
	call := conn.BusObject().Call("org.freedesktop.DBus.AddMatch", 0,
		"interface='"+notificationDBusInterface+"',member='"+notificationDBusMember+"',type='method_call',eavesdrop='true'")
	if call.Err != nil {
		return NotificationReceiver{}, call.Err
	}

	c := make(chan *dbus.Message, 10)
	conn.Eavesdrop(c)
	return NotificationReceiver{channel: c, connection: conn, systemRepeatsTwice: systemRepeatsTwice}, nil
}

// This method closes the `NotificationReceiver` object.
func (notificationReceiver NotificationReceiver) Close() {
	notificationReceiver.connection.Close()
	close(notificationReceiver.channel)
}

func handleUnmarshallingForChannel(notificationReceiver NotificationReceiver, outputChannel chan Notification) {
	for v := range notificationReceiver.channel {
		outputChannel <- unmarshallNotification(v)
		if notificationReceiver.systemRepeatsTwice {
			<-notificationReceiver.channel
		}
	}
}

// This method returns a channel that receives `Notification` objects.
func (notificationReceiver NotificationReceiver) GetChannel() chan Notification {
	outputChannel := make(chan Notification, 10)
	go handleUnmarshallingForChannel(notificationReceiver, outputChannel)
	return outputChannel
}

func unmarshallNotification(dbusMsg *dbus.Message) Notification {
	if len(dbusMsg.Body) >= 8 {
		return Notification{Body: NotificationBody{
			ApplicationName:   dbusMsg.Body[0].(string),
			ReplacesID:        dbusMsg.Body[1].(uint32),
			NotificationIcon:  dbusMsg.Body[2].(string),
			Summary:           dbusMsg.Body[3].(string),
			Body:              dbusMsg.Body[4].(string),
			Actions:           dbusMsg.Body[5].([]string),
			Hints:             dbusMsg.Body[6].(map[string]dbus.Variant),
			ExpirationTimeout: dbusMsg.Body[7].(int32),
		}, Error: nil}
	}
	return Notification{Error: errors.New("notifications: index out of range; DBus message body is not long enough")}

}

// This method blocks until a notification is received and returns the subsequent `Notification` object.
func (notificationReceiver NotificationReceiver) GetBlocking() Notification {
	message := <-notificationReceiver.channel
	if notificationReceiver.systemRepeatsTwice {
		<-notificationReceiver.channel
	}
	return unmarshallNotification(message)
}
