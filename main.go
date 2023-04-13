package main

import (
  "net/http"

  "github.com/gin-gonic/gin"
  "github.com/joshnavdev/loans-server/config"
  "github.com/joshnavdev/loans-server/routes"
)


func main() {
  envs := config.GetEnv()

  r := gin.Default()

  r.GET("/", func(ctx *gin.Context) {
    ctx.JSON(200, gin.H{"message": "Hello WOrld"})
  })
  
  routes.LoadRoutes(r)
  
  r.NoRoute(func(ctx *gin.Context) {
    ctx.JSON(http.StatusNotFound, gin.H{"message": "Not Found"})
  })
  
  r.Run(envs.Application.Addr)
}
