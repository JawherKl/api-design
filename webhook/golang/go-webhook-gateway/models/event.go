package models

import "time"

type WebhookEvent struct {
	Source     string                 `bson:"source"`
	ReceivedAt time.Time              `bson:"received_at"`
	Headers    map[string]string      `bson:"headers"`
	Payload    map[string]interface{} `bson:"payload"`
}
