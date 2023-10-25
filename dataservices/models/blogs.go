package models

import (
	"database/sql"
)

type Blog struct {
	ID          sql.NullString `json:"type:uuid;primary_key;"`
	Title       sql.NullString `json:"title"`
	Description sql.NullString `json:"description"`
	CreatedAt   sql.NullTime   `json:"created_at"`
	UpdatedAt   sql.NullTime   `json:"updated_at"`
	DeletedAt   sql.NullTime   `json:"deleted_at"`
}
