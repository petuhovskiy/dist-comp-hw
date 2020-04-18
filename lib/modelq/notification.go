package modelq

type Notification struct {
	// Type of notification, "sms" or "email"
	Type string

	// Whom send message to
	Recipient string

	// Content of the notification
	Content string
}
