package repository

import (
	"gorm.io/gorm"

	"github.com/zorth44/zorth-blog-server/internal/model"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

// 在这里添加数据库操作方法
// 例如:
// func (r *Repository) GetUsers() ([]model.User, error) {
//     var users []model.User
//     result := r.db.Find(&users)
//     return users, result.Error
// }

func (r *Repository) GetPosts() ([]model.Post, error) {
	var posts []model.Post
	result := r.db.Find(&posts)
	return posts, result.Error
}

func (r *Repository) GetPostBySlug(slug string) (*model.Post, error) {
	var post model.Post
	result := r.db.Where("slug = ?", slug).First(&post)
	if result.Error != nil {
		return nil, result.Error
	}
	return &post, nil
}
