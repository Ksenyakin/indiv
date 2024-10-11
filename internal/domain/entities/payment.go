// internal/domain/entities/payment.go

package entities

import "time"

type Payment struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
	Status    string    `json:"status"` // "PENDING", "COMPLETED", "FAILED"
}
