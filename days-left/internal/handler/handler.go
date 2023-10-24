package handler

import (
	"days-left/internal/middleware"
	"days-left/internal/service"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	router.Use(middleware.TokenMiddlerware())

	router.Use(gin.Recovery())

	api := router.Group("/api")
	{
		v1 := api.Group("/v1")
		{
			days := v1.Group("/days")
			{
				days.GET("/", h.GetDays)
			}
		}
	}

	return router
}
