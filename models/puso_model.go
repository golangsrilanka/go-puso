package models

import "gorm.io/gorm"

type Puso struct {
	gorm.Model `swaggerignore:"true"`
	Name       string  `json:"name,omitempty"`
	Color      string  `json:"color,omitempty"`
	Weight     float32 `json:"weight,omitempty"`
	Owner      string  `json:"owner,omitempty"`
	Laziness   int     `json:"laziness,omitempty"`
}
