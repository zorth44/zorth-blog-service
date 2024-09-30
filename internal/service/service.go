package service

import (
	"github.com/zorth44/zorth-blog-server/internal/model"
	"github.com/zorth44/zorth-blog-server/internal/repository"
)

type Service struct {
	repo *repository.Repository
}

func NewService(repo *repository.Repository) *Service {
	return &Service{
		repo: repo,
	}
}

// 在这里添加服务方法
// 例如:
// func (s *Service) GetUsers() ([]model.User, error) {
//     return s.repo.GetUsers()
// }

func (s *Service) GetPosts() ([]model.Post, error) {
	return s.repo.GetPosts()
}

func (s *Service) GetPostBySlug(slug string) (*model.Post, error) {
	return s.repo.GetPostBySlug(slug)
}
