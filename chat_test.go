package slack

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"testing"
)

func TestPostMessage(t *testing.T) {
	t.Skip("goes over the network")
	client := New(os.Getenv("SLACK_TOKEN"))
	users, err := client.Users.List(context.Background(), nil)
	if err != nil {
		t.Fatal(err)
	}
	var kevinuser string
	for _, u := range users.Members {
		if u.Name == "kevin" {
			kevinuser = u.ID
		}
	}
	resp, err := client.Chat.PostMessage(context.Background(), url.Values{
		"channel": []string{kevinuser},
		"text":    []string{"Hello from meterbot"},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Printf("resp: %#v\n", resp)
}
