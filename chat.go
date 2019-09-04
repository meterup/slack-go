package slack

import (
	"context"
	"net/url"
)

type ChatService struct {
	client *Client
}

type response struct {
	OK      bool      `json:"ok"`
	Channel string    `json:"channel"`
	TS      SlackTime `json:"ts"`
	Error   string    `json:"error"`
	Warning string    `json:"warning"`
}

// https://api.slack.com/methods/chat.postMessage
func (c *ChatService) PostMessage(ctx context.Context, data url.Values) (*response, error) {
	resp := new(response)
	err := c.client.CreateResource(ctx, "/api/chat.postMessage", data, resp)
	return resp, err
}
