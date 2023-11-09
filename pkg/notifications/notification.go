package notifications

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"sync"

	"github.com/kintuda/tech-challenge-pismo/pkg/queue"
	"github.com/rs/zerolog/log"
)

var _ TransactionNotification = (*WebhookNotification)(nil)
var _ TransactionNotification = (*SMSNotification)(nil)

type TransactionNotification interface {
	Trigger(ctx context.Context, trx interface{}, wg *sync.WaitGroup) error
}

type SMSNotification struct{}

type NotificationOrchestrator struct {
	Integrations []TransactionNotification
	Queue        queue.InMemoryQueue
}

func NewNotificationOrchestrator() *NotificationOrchestrator {
	return &NotificationOrchestrator{
		Integrations: []TransactionNotification{
			NewWebhookNotification(),
			&SMSNotification{},
		},
	}
}

func (n *NotificationOrchestrator) Read() {
	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}

func (n *NotificationOrchestrator) SendNotifications(payload interface{}) error {
	ctx := context.Background()
	wg := sync.WaitGroup{}

	for _, in := range n.Integrations {
		wg.Add(1)
		go in.Trigger(ctx, payload, &wg)
	}

	wg.Wait()

	return nil
}

func (*SMSNotification) Trigger(ctx context.Context, payload interface{}, wg *sync.WaitGroup) error {
	defer wg.Done()
	log.Info().Msg("SMS sent")
	return nil
}

type WebhookNotification struct {
	URL    string
	Client *http.Client
}

func NewWebhookNotification() *WebhookNotification {
	return &WebhookNotification{
		URL:    "https://webhook.site/3a687e26-d497-4c33-bc54-5699ace6de95",
		Client: http.DefaultClient,
	}
}

func (w *WebhookNotification) Trigger(ctx context.Context, payload interface{}, wg *sync.WaitGroup) error {
	content, err := json.Marshal(payload)

	if err != nil {
		log.Error().Err(err).Msg("error while sending notification")
		return err
	}

	req, err := http.NewRequestWithContext(ctx, http.MethodPost, w.URL, bytes.NewReader(content))

	if err != nil {
		log.Error().Err(err).Msg("error while building request ")
		return err
	}

	res, err := w.Client.Do(req)

	if err != nil {
		log.Error().Err(err).Msg("error while sending request")
		return err
	}

	if res.StatusCode != 200 && res.StatusCode != 201 {
		return errors.New("request failed with non sucessfull request")
	}

	respBody, err := io.ReadAll(res.Body)

	if err != nil {
		log.Error().Err(err).Msg("error while reading body")
		return err
	}

	log.Info().Msg(string(respBody))

	defer func() {
		if err := res.Body.Close(); err != nil {
			log.Error().Err(err).Msg("error while closing body stream")
		}

		wg.Done()
	}()

	return nil
}
