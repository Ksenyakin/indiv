// internal/infrastructure/adapters/notification_adapter.go

package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"indiv/internal/domain/entities"
)

type NotificationAdapter interface {
	SendNotification(ctx context.Context, notification *entities.Notification) error
}

type notificationAdapter struct {
	client      *http.Client
	endpointURL string
	apiKey      string
}

func NewNotificationAdapter(endpointURL, apiKey string) NotificationAdapter {
	return &notificationAdapter{
		client: &http.Client{
			Timeout: 5 * time.Second,
		},
		endpointURL: endpointURL,
		apiKey:      apiKey,
	}
}

func (n *notificationAdapter) SendNotification(ctx context.Context, notification *entities.Notification) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"user_id": notification.UserID,
		"message": notification.Message,
	})
	if err != nil {
		return fmt.Errorf("error marshalling notification request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", n.endpointURL+"/notifications", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating notification request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+n.apiKey)

	resp, err := n.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending notification request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("notification request failed: %s", string(bodyBytes))
	}

	// Обновление статуса уведомления
	notification.CreatedAt = time.Now()
	notification.Read = false

	return nil
}
