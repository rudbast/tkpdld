package slack

import (
	"context"
	"errors"

	"github.com/rudbast/tkpdld/config"

	slackapi "github.com/nlopes/slack"
)

var (
	// Singleton.
	client SlackClientInterface
)

type (
	SlackClientInterface interface {
		New(ctx context.Context) error
		UpdateStatus(ctx context.Context, status, emoji string) error
	}

	SlackAPIClient struct {
		client *slackapi.Client
	}
)

func (sc *SlackAPIClient) New(ctx context.Context) error {
	cfg := config.Get()
	sc.client = slackapi.New(cfg.SlackAPI.Token)
	return nil
}

func (sc *SlackAPIClient) UpdateStatus(ctx context.Context, status, emoji string) error {
	if sc.client == nil {
		return errors.New("uninitialized client")
	}

	return sc.client.SetUserCustomStatusContext(ctx, status, emoji)
}

func Load(ctx context.Context) error {
	slackAPIClient := &SlackAPIClient{}
	err := slackAPIClient.New(ctx)
	if err != nil {
		return err
	}

	client = slackAPIClient
	return nil
}

func Get() SlackClientInterface {
	return client
}
