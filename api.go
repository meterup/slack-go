package slack

import (
	"context"
	"net/url"
)

// Convenience endpoints

// SendMessage sends the text to the given channel. For more granular control
// use the client.Chat.PostMessage endpoint.
func (c *Client) SendMessage(channel string, text string) error {
	data := url.Values{}
	data.Set("channel", channel)
	data.Set("text", text)
	_, err := c.Chat.PostMessage(context.Background(), data)
	return err
}
