package models

import "time"

type Product struct {
	ID          int
	Description string
	ProductType string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}
