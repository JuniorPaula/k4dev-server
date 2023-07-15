package models

import (
	"errors"
	"sort"
)

type Category struct {
	ID       int64      `json:"id,omitempty"`
	Name     string     `json:"name,omitempty"`
	ParentID int64      `json:"parent_id,omitempty"`
	Path     string     `json:"path,omitempty"`
	Children []Category `json:"children,omitempty"`
}

// HanlderCategory validates the category fields
func (c *Category) HanlderCategory() error {
	if c.Name == "" {
		return errors.New("name is required")
	}

	return nil
}

// WithPath returns a slice of categories with the path field filled
func (c *Category) WithPath(categories []Category) []Category {
	getParent := func(categories []Category, parentID int64) *Category {
		for _, parent := range categories {
			if parent.ID == parentID {
				return &parent
			}
		}
		return nil
	}

	categoriesWithPath := make([]Category, len(categories))
	copy(categoriesWithPath, categories)

	for i := range categoriesWithPath {
		path := categoriesWithPath[i].Name
		parent := getParent(categoriesWithPath, categoriesWithPath[i].ParentID)

		for parent != nil {
			path = parent.Name + " > " + path
			parent = getParent(categoriesWithPath, parent.ParentID)
		}

		categoriesWithPath[i].Path = path
	}

	// sort by path
	sort.Slice(categoriesWithPath, func(i, j int) bool {
		return categoriesWithPath[i].Path < categoriesWithPath[j].Path
	})

	return categoriesWithPath
}

// ToTree returns a slice of categories in tree format
func (c *Category) ToTree(categories []Category, tree []Category) []Category {
	if tree == nil {
		tree = make([]Category, 0)
		for _, c := range categories {
			if c.ParentID == 0 {
				tree = append(tree, c)
			}
		}
	}

	for i := range tree {
		parentNode := &tree[i]
		parentNode.Children = c.ToTree(categories, getChildren(categories, parentNode.ID))
	}

	return tree
}

// getChildren returns a slice of categories that have the given parentID
func getChildren(categories []Category, parentID int64) []Category {
	children := make([]Category, 0)
	for _, c := range categories {
		if c.ParentID == parentID {
			children = append(children, c)
		}
	}
	return children
}
