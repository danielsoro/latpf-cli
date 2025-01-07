package client

import (
	"github.com/sogko/go-wordpress"
	"log/slog"
)

var _ Wordpress = (*WithViper)(nil)

type WithViper struct {
	client *wordpress.Client
}

func (w WithViper) CreatePost(title, content, status string) (*wordpress.Post, error) {
	post, _, _, err := w.client.Posts().Create(&wordpress.Post{
		Title: wordpress.Title{
			Raw: title,
		},
		Content: wordpress.Content{
			Raw: content,
		},
		Status: status,
	})

	if err != nil {
		return nil, err
	}

	return post, nil
}

func (w WithViper) GetPosts() []wordpress.Post {
	posts, _, _, err := w.client.Posts().List(nil)
	if err != nil {
		slog.Error("Error getting posts", slog.String("error", err.Error()))
	}

	return posts
}
