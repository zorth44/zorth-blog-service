package main

import (
	"log"

	"github.com/zorth44/zorth-blog-server/config"
	"github.com/zorth44/zorth-blog-server/internal/handler"
	"github.com/zorth44/zorth-blog-server/internal/repository"
	"github.com/zorth44/zorth-blog-server/internal/service"
	"github.com/zorth44/zorth-blog-server/pkg/database"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	// 加载配置
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("加载配置失败: %v", err)
	}

	// 初始化数据库
	db, err := database.InitDB(cfg.Database)
	if err != nil {
		log.Fatalf("初始化数据库失败: %v", err)
	}

	// 设置依赖
	repo := repository.NewRepository(db)
	svc := service.NewService(repo)
	h := handler.NewHandler(svc)

	// 设置Gin路由
	r := gin.Default()
	r.Use(cors.Default())
	h.SetupRoutes(r)

	// 启动服务器
	if err := r.Run(cfg.Server.Address); err != nil {
		log.Fatalf("启动服务器失败: %v", err)
	}
}
