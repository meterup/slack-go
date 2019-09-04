package slack

import (
	"context"
	"net/url"
)

type UserService struct {
	client *Client
}

type User struct {
	ID      string    `json:"id"`
	Name    string    `json:"name"`
	Updated SlackTime `json:"updated"`
}

type UsersListResponse struct {
	Members []*User `json:"members"`
}

func (u *UserService) List(ctx context.Context, data url.Values) (*UsersListResponse, error) {
	resp := new(UsersListResponse)
	err := u.client.ListResource(ctx, "/api/users.list", data, resp)
	return resp, err
}
