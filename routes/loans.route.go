package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoanRoute(r *gin.Engine) {
  loanRouter := r.Group("/loans")

  loanRouter.GET("/", func (c *gin.Context)  {
    c.JSON(http.StatusOK,  "loans")
  })
}
