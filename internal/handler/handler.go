package handler

import (
	"net/http"

	"github.com/zorth44/zorth-blog-server/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(svc *service.Service) *Handler {
	return &Handler{service: svc}
}

func (h *Handler) SetupRoutes(r *gin.Engine) {
	// 在这里设置路由
	// 例如: r.GET("/users", h.GetUsers)
	api := r.Group("/api")
	{
		api.GET("/posts", h.GetPosts)
	}
}

// 在这里添加处理函数
// 例如: func (h *Handler) GetUsers(c *gin.Context) { ... }

func (h *Handler) GetPosts(c *gin.Context) {
	posts, err := h.service.GetPosts()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "获取文章失败"})
		return
	}
	c.JSON(http.StatusOK, posts)
}
