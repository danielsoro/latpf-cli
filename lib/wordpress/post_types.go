package wordpress

import "errors"

type PostType string

func (e *PostType) String() string {
	return string(*e)
}

func (e *PostType) Set(v string) error {
	switch v {
	case "public", "premium", "curation":
		*e = PostType(v)
		return nil
	default:
		return errors.New(`must be one of "public", "premium", or "curation"`)
	}
}

func (e *PostType) Type() string {
	return "postType"
}
