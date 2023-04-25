package service

import (
	common "github.com/joshnavdev/loans-server/commons"
	request "github.com/joshnavdev/loans-server/commons/requests"
)

type LoanService struct {}

func NewLoanService() LoanService {
  return LoanService{}
}

func (service *LoanService) Simulate(loanRequest request.SimulateLoanRequest) common.Loan {
  loan := common.NewLoan(loanRequest)

  loan.Simulate()

  return loan
}

