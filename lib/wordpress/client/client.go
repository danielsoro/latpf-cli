package client

import (
	"errors"
	"github.com/sogko/go-wordpress"
	"github.com/spf13/viper"
)

type WordPressClientType int

const (
	VIPER WordPressClientType = iota + 1
)

type Wordpress interface {
	GetPosts() []wordpress.Post
	CreatePost(title, content, status string) (*wordpress.Post, error)
}

func NewWordPressClient(clientType WordPressClientType) (Wordpress, error) {
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
