package main

import (
	"log"
	"os"

	"forma/internal/cache"
	"forma/internal/config"
	"forma/internal/handler"
	"forma/internal/middleware"
	repository "forma/internal/repostory"
	"forma/internal/service"
	loadenv "forma/pkg/loadEnv"

	"github.com/gin-gonic/gin"
)

func main() {
	loadenv.LoadEnv()

	postgres, err := config.NewPostgresDB()
	if err != nil {
		log.Fatal("postgres ulanish xatosi: ", err)
	}

	redis, err := config.NewRedisClient()
	if err != nil {
		log.Fatal("redis ulanish xatosi: ", err)
	}

	leadCache := cache.NewLeadCache(redis)
	leadRepo := repository.NewLeadRepository(postgres)
	leadService := service.NewLeadService(leadRepo, leadCache)
	leadHandler := handler.NewLeadHandler(leadService)

	r := gin.Default()

	// public
	r.POST("/leads", leadHandler.Create)

	// protected
	admin := r.Group("/")
	admin.Use(middleware.AdminAuthMiddleware())
	{
		admin.GET("/leads", leadHandler.GetAll)
		admin.GET("/leads/:id", leadHandler.GetByID)
		admin.DELETE("/leads/:id", leadHandler.Delete)
	}


	log.Println("server ishga tushdi: :8080")
	if err := r.Run(os.Getenv("PORT")); err != nil {
		log.Fatal("server run xatosi: ", err)
	}
}