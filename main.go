package main

import (
	"go-project/controllers"
	"go-project/middlewares"
	"go-project/view"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	r.HTMLRender = &TemplRender{}

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "", view.Home("User"))
	})

	public := r.Group("/api")
	//authentication
	public.POST("/register", controllers.Register)
	public.POST("/login", controllers.Login)

	protected := r.Group("/api/admin")
	protected.Use(middlewares.JwtAuthMiddleware())
	protected.GET("/user", controllers.CurrentUser)

	r.Run(":8080")

}
