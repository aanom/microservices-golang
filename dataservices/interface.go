package dataservices

import (
	"database/sql"

	"github.com/backend-services/dataservices/models"
)

// PostgresClient stored reference to the DB
type PostgresClient struct {
	DB          *sql.DB
	DBRadiuzDWH *sql.DB
	DBYor24DWH  *sql.DB
}

// IPostgresClient - DB interface
type IPostgresClient interface {
	Connect()
	GetBlogs() ([]models.Blog, error)
	GetBlogByID(ID string) (models.Blog, error)
	DeleteBlogByID(ID string) error
	UpdateBlogByID(ID, title, description string) error
}
