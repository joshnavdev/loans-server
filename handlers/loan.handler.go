package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joshnavdev/loans-server/commons/requests"
	mapper "github.com/joshnavdev/loans-server/mappers"
	service "github.com/joshnavdev/loans-server/services"
)

type LoanHandler struct {
  loanService service.LoanService
}

func NewLoanHandler(loanService service.LoanService) LoanHandler {
  return LoanHandler{
    loanService: loanService,
  }
}


func (handler *LoanHandler) Simulate(ctx *gin.Context) {
  var simulateLoanRequest request.SimulateLoanRequest

  err := ctx.ShouldBindJSON(&simulateLoanRequest)
  if err != nil {
    ctx.AbortWithError(http.StatusBadRequest, err)
    return
  }

  simulation := handler.loanService.Simulate(simulateLoanRequest)

  response := mapper.LoanSimulationToDTO(simulation)
  
  ctx.JSON(http.StatusOK, response)
}
