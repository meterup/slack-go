package slack

import (
	"context"
	"net/url"
)

type UserService struct {
	client *Client
}

type User struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Updated SlackTime    `json:"updated"`
	Profile *UserProfile `json:"profile"`
}

type UserResponse struct {
	User *User `json:"user"`
}

type UserProfile struct {
	Email string `json:"email"`
	Team  string `json:"team"`
}

type UsersListResponse struct {
	Members []*User `json:"members"`
}

func (u *UserService) List(ctx context.Context, data url.Values) (*UsersListResponse, error) {
	resp := new(UsersListResponse)
	err := u.client.ListResource(ctx, "/api/users.list", data, resp)
	return resp, err
}

// Get information about a particular user.
// https://api.slack.com/methods/users.info
func (u *UserService) Get(ctx context.Context, data url.Values) (*User, error) {
	resp := new(UserResponse)
	err := u.client.ListResource(ctx, "/api/users.info", data, resp)
	if err != nil {
		return nil, err
	}
	return resp.User, nil
}
