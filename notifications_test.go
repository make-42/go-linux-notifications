package notifications

import (
	"testing"

	"github.com/godbus/dbus/v5"
	"github.com/stretchr/testify/assert"
)

func TestNewNotificationReceiver(t *testing.T) {
	_, err := NewNotificationReceiver(true)
	if err != nil {
		t.Errorf("%q", err)
	}
}

func TestUnmarshallNotification(t *testing.T) {
	expectedUnmarshalledNotification := Notification{Body: NotificationBody{
		ApplicationName:  "application",
		ReplacesID:       uint32(0),
		NotificationIcon: "file://test.test",
		Summary:          "testsummary",
		Body:             "testbody",
		Actions:          []string{"testaction1", "testaction2", "testaction3"},
		Hints: map[string]dbus.Variant{
			"testhint1": dbus.MakeVariant(123456789),
		},
		ExpirationTimeout: int32(1),
	}, Error: nil}
	testMessageBody := []interface{}{
		expectedUnmarshalledNotification.Body.ApplicationName,
		expectedUnmarshalledNotification.Body.ReplacesID,
		expectedUnmarshalledNotification.Body.NotificationIcon,
		expectedUnmarshalledNotification.Body.Summary,
		expectedUnmarshalledNotification.Body.Body,
		expectedUnmarshalledNotification.Body.Actions,
		expectedUnmarshalledNotification.Body.Hints,
		expectedUnmarshalledNotification.Body.ExpirationTimeout,
	}
	testMesssage := &dbus.Message{Body: testMessageBody}
	notification := UnmarshallNotification(testMesssage)
	if notification.Error != nil {
		t.Errorf("%q\n", notification.Error)
	}
	assert.Equal(t, notification, expectedUnmarshalledNotification, "error: unmarshalled notification doesn't match expected result")
}
