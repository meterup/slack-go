package slack

import (
	"context"
	"net/url"
	"os"
	"testing"
)

func TestUsersGet(t *testing.T) {
	t.Skip("goes over the network")
	client := New(os.Getenv("SLACK_TOKEN"))
	user, err := client.Users.Get(context.Background(), url.Values{
		"user": []string{"UKPDARH39"},
	})
	if err != nil {
		t.Fatal(err)
	}
	if user.Profile.Email != "kevin@meter.com" {
		t.Errorf("should have parsed email address, did not")
	}
}
