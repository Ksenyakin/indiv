// internal/infrastructure/adapters/payment_adapter.go

package adapters

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"indiv/internal/domain/entities"
)

type PaymentAdapter interface {
	ProcessPayment(ctx context.Context, payment *entities.Payment) error
}

type paymentAdapter struct {
	client      *http.Client
	endpointURL string
	apiKey      string
}

func NewPaymentAdapter(endpointURL, apiKey string) PaymentAdapter {
	return &paymentAdapter{
		client: &http.Client{
			Timeout: 10 * time.Second,
		},
		endpointURL: endpointURL,
		apiKey:      apiKey,
	}
}

func (p *paymentAdapter) ProcessPayment(ctx context.Context, payment *entities.Payment) error {
	requestBody, err := json.Marshal(map[string]interface{}{
		"user_id": payment.UserID,
		"amount":  payment.Amount,
	})
	if err != nil {
		return fmt.Errorf("error marshalling payment request: %w", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", p.endpointURL+"/payments", bytes.NewBuffer(requestBody))
	if err != nil {
		return fmt.Errorf("error creating payment request: %w", err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+p.apiKey)

	resp, err := p.client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending payment request: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		return fmt.Errorf("payment request failed: %s", string(bodyBytes))
	}

	var response struct {
		Success bool `json:"success"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return fmt.Errorf("error decoding payment response: %w", err)
	}

	if !response.Success {
		return errors.New("payment failed")
	}

	// Обновление статуса платежа
	payment.Status = "COMPLETED"
	payment.Timestamp = time.Now()

	return nil
}
