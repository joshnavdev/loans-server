package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshnavdev/loans-server/models"
	service "github.com/joshnavdev/loans-server/services"
)

type AuthHandler struct {
  authService service.AuthService
}

func NewAuthHandler(authService service.AuthService) AuthHandler {
  return AuthHandler{
    authService: authService,
  }
}

func (handler *AuthHandler) SignUp(ctx *gin.Context) {
  var user models.User

  if err := ctx.ShouldBindJSON(&user); err != nil {
    ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
    return
  }

  // validar input

  if err := handler.authService.SignUp(user); err != nil {
    ctx.JSON(http.StatusInternalServerError, gin.H{"errorCode": "AUTH001", "message": "Internal Server Error", "error": err.Error()})
    return
  }
  
  ctx.JSON(http.StatusNoContent, nil)
}
