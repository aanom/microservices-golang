package dataservices

import (
	"database/sql"
	"errors"
	"time"

	"github.com/sirupsen/logrus"

	"github.com/backend-services/dataservices/models"
)

// GetBlogByID - Function get blog by id
// params @ID
func (pc *PostgresClient) GetBlogByID(ID string) (blog models.Blog, err error) {
	logHeader := "Func[GetBlogByID] : @dataservices"
	sqlQuery := "select id, title, description, created_at, updated_at FROM blogs WHERE id = $1 AND deleted_at = $2"
	err = pc.DB.QueryRow(sqlQuery, ID, nil).Scan(&blog.ID, &blog.Title, &blog.Description, &blog.CreatedAt, &blog.UpdatedAt)
	if err != nil {
		logrus.Error(logHeader, err.Error())
		return
	}
	return
}

// DeleteBlogByID - Function for Soft delete blog
// params @ID
func (pc *PostgresClient) DeleteBlogByID(ID string) (err error) {
	logHeader := "Func[DeleteBlogByID] : @dataservices"
	var id sql.NullString
	sqlQuery := `UPDATE blogs SET deleted_at = $1 WHERE ID = $2 returning id;`
	err = pc.DB.QueryRow(sqlQuery, time.Now(), ID).Scan(&id)
	if err != nil {
		logrus.Error(logHeader, err.Error())
		return
	}
	if !id.Valid {
		logrus.Error(logHeader, errors.New("invalid id recived"))
		return
	}
	return
}

// UpdateBlogByID - Function for Update blog By ID
// params @ID
func (pc *PostgresClient) UpdateBlogByID(ID, title, description string) (err error) {
	logHeader := "@dataservices Func[InsertUser] : "
	var id sql.NullString

	sqlQuery := `UPDATE blogs SET title = $1, description = $2 WHERE id = $3 returning id;`
	err = pc.DB.QueryRow(sqlQuery, title, description).Scan(&id)
	if err != nil {
		logrus.Error(logHeader, err)
		return
	}
	if !id.Valid {
		logrus.Error(logHeader, errors.New("invalid id recived"))
		return
	}
	return
}

// GetBlogs : Fetch all the blogs
func (pc *PostgresClient) GetBlogs() (blogs []models.Blog, err error) {
	query := `select id, title, description, created_at, updated_at FROM blogs`
	rows, err := pc.DB.Query(query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var blog models.Blog
		err = rows.Scan(
			&blog.ID,
			&blog.Title,
			&blog.Description,
			&blog.CreatedAt,
			&blog.UpdatedAt,
		)
		if err != nil && err != sql.ErrNoRows {
			return
		}
		blogs = append(blogs, blog)
	}
	return
}
