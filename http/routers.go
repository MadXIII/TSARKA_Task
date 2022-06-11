package http

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/madxiii/tsarka_task/service"
)

type API struct {
	services *service.Service
}

func NewAPI(services service.Service) *API {
	return &API{services: &services}
}

func (h *API) InitRoutes() http.Handler {
	router := gin.New()

	rest := router.Group("/rest")
	{
		rest.POST("/substr/find", h.Find)
		rest.POST("/email/check", h.EmailSearch)

		counter := rest.Group("/counter")
		{
			counter.GET("/val", h.CountGet)
			counter.POST("/add/:number", h.CountAdd)
			counter.POST("/sub/:number", h.CountSub)
		}
		user := rest.Group("/user")
		{
			user.GET("/:id", h.GetUser)
			user.POST("", h.CreateUser)
			user.PUT("/:id", h.UpdateUser)
			user.DELETE("/:id", h.DeleteUser)
		}
		hash := rest.Group("/hash")
		{
			hash.POST("/calc", h.Calculate)
			hash.GET("/result/:id", h.GetResult)
		}
	}

	return router
}
