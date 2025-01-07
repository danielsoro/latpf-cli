package wordpress

import (
	"errors"
	"log/slog"

	"github.com/sogko/go-wordpress"
	"github.com/spf13/viper"
)

type ClientType int

const (
	VIPER ClientType = iota + 1
)

type Wordpress interface {
	GetPosts() []wordpress.Post
	CreatePost(title, content, status string) (*wordpress.Post, error)
}

var _ Wordpress = (*WithViper)(nil)

type WithViper struct {
	client *wordpress.Client
}

func NewWordPressClient(clientType ClientType) (Wordpress, error) {
	switch clientType {
	case VIPER:
		return WithViper{
			client: wordpress.NewClient(&wordpress.Options{
				BaseAPIURL: viper.GetString("url"),
				Username:   viper.GetString("credentials.username"),
				Password:   viper.GetString("credentials.password"),
			}),
		}, nil
	default:
		return nil, errors.New("type is required to define the client")
	}
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
