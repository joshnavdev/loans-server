package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshnavdev/loans-server/handlers"
	repository "github.com/joshnavdev/loans-server/repositories"
	service "github.com/joshnavdev/loans-server/services"
)

func AuthRoute(r *gin.Engine) {
  authRouter := r.Group("/auth")

  userRepository := repository.NewUserRepository()
  authService := service.NewAuthService(userRepository)
  authHandler := handlers.NewAuthHandler(authService)

  authRouter.POST("/signup", authHandler.SignUp)
}
