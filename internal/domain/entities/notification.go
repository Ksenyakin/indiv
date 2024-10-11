// internal/domain/entities/notification.go

package entities

import "time"

type Notification struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	Read      bool      `json:"read"`
}
