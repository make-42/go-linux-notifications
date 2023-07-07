package linuxnotifications

import "github.com/godbus/dbus/v5"

// See Table 1 @ https://specifications.freedesktop.org/notification-spec/notification-spec-latest.html
type NotificationBody struct {
	ApplicationName   string                  // This is the optional name of the application sending the notification. This should be the application's formal name, rather than some sort of ID. An example would be "FredApp E-Mail Client," rather than "fredapp-email-client."
	ReplacesID        uint32                  // An optional ID of an existing notification that this notification is intended to replace.
	NotificationIcon  string                  // The notification icon. See Icons and Images Formats.
	Summary           string                  // This is a single line overview of the notification. For instance, "You have mail" or "A friend has come online". It should generally not be longer than 40 characters, though this is not a requirement, and server implementations should word wrap if necessary. The summary must be encoded using UTF-8.
	Body              string                  // This is a multi-line body of text. Each line is a paragraph, server implementations are free to word wrap them as they see fit. The body may contain simple markup as specified in Markup. It must be encoded using UTF-8. If the body is omitted, just the summary is displayed.
	Actions           []string                // The actions send a request message back to the notification client when invoked. This functionality may not be implemented by the notification server, conforming clients should check if it is available before using it (see the GetCapabilities message in Protocol). An implementation is free to ignore any requested by the client. As an example one possible rendering of actions would be as buttons in the notification popup. Actions are sent over as a list of pairs. Each even element in the list (starting at index 0) represents the identifier for the action. Each odd element in the list is the localized string that will be displayed to the user. The default action (usually invoked by clicking the notification) should have a key named "default". The name can be anything, though implementations are free not to display it.
	Hints             map[string]dbus.Variant // Hints are a way to provide extra data to a notification server that the server may be able to make use of. See Hints for a list of available hints.
	ExpirationTimeout int32                   // The timeout time in milliseconds since the display of the notification at which the notification should automatically close. If -1, the notification's expiration time is dependent on the notification server's settings, and may vary for the type of notification. If 0, the notification never expires.
}

type Notification struct {
	Body  NotificationBody
	Error error
}

type NotificationReceiver struct {
	channel            chan *dbus.Message
	connection         *dbus.Conn
	systemRepeatsTwice bool
}
