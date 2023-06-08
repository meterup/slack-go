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

type permalinkResponse struct {
	OK        bool   `json:"ok"`
	Channel   string `json:"channel"`
	Permalink string `json:"permalink"`
	Error     string `json:"error"`
}

// https://api.slack.com/methods/chat.postMessage
func (c *ChatService) PostMessage(ctx context.Context, data url.Values) (*response, error) {
	resp := new(response)
	err := c.client.CreateResource(ctx, "/api/chat.postMessage", data, resp)
	return resp, err
}

// https://api.slack.com/methods/chat.delete
func (c *ChatService) DeleteMessage(ctx context.Context, channelID, messageID string) (*response, error) {
	data := url.Values{
		"channel": []string{channelID},
		"ts":      []string{messageID},
	}
	resp := new(response)
	err := c.client.CreateResource(ctx, "/api/chat.delete", data, resp) // Delete is a POST
	return resp, err
}

// https://api.slack.com/methods/chat.getPermalink
//
// `channelID` has to be the ID of the channel (C*), this endpoint does not
// support the encoded channel names (#foo).
func (c *ChatService) GetPermalink(ctx context.Context, channelID, messageID string) (*permalinkResponse, error) {
	data := url.Values{
		"channel":    []string{channelID},
		"message_ts": []string{messageID},
	}
	resp := new(permalinkResponse)
	err := c.client.CreateResource(ctx, "/api/chat.getPermalink", data, resp)
	return resp, err
}
