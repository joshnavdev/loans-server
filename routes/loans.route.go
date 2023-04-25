package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/joshnavdev/loans-server/handlers"
	service "github.com/joshnavdev/loans-server/services"
)

func LoanRoute(r *gin.Engine) {
  loanRouter := r.Group("/loans")

  loanService := service.NewLoanService()
  loanHandler := handlers.NewLoanHandler(loanService)

  loanRouter.POST("/simulate", loanHandler.Simulate)
}
