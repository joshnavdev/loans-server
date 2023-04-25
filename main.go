package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshnavdev/loans-server/config"
	middleware "github.com/joshnavdev/loans-server/middlewares"
	"github.com/joshnavdev/loans-server/routes"
)


func main() {
  envs := config.GetEnv()

  r := gin.Default()

  // r.Use(gin.CustomRecovery(middleware.ErrorHandler))
  r.Use(middleware.ErrorHandler)

  r.GET("/", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Hello WOrld"})
  })
  
  routes.LoadRoutes(r)
  
  r.NoRoute(func(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
  })
  
  r.Run(envs.Application.Addr)
}
