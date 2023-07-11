package models

import "errors"

type Category struct {
	ID       int64  `json:"id,omitempty"`
	Name     string `json:"name,omitempty"`
	ParentID int64  `json:"parent_id,omitempty"`
}

func (c *Category) HanlderCategory() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	return nil
}
