package middleware

import (
	"errors"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func getErrorMessage(fe validator.FieldError) string {
  fmt.Println(fe)
  switch fe.Tag() {
    case "required":
      return "This field es required"
  }

  return "Unknow error"
}

type ErrorMessage struct {
  Field string `json:"field"`
  Message string `json:"message"`
}

func getErrorResponse(err error) []ErrorMessage {
  var ve validator.ValidationErrors
  var errorResponse []ErrorMessage

  if errors.As(err, &ve) {
    for _, fieldError := range ve {
      newErrorMessage := ErrorMessage{fieldError.Field(), getErrorMessage(fieldError)}
      errorResponse = append(errorResponse, newErrorMessage)
    }
  }

  return errorResponse
}

func ErrorHandler(ctx *gin.Context) {
  ctx.Header("Content-Type", "application/json")
  ctx.Next()

  // TODO: Validar otros tipos de errores
  if ctx.Errors != nil {
    for _, err := range ctx.Errors {
      ctx.JSON(-1, gin.H{"error": getErrorResponse(err)})
    }
  }
}
