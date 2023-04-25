package mapper

import (
	"github.com/google/uuid"
	common "github.com/joshnavdev/loans-server/commons"
	"github.com/joshnavdev/loans-server/dtos"
)

func LoanSimulationToDTO(loan common.Loan) dtos.LoanSimulationDTO {
  var loanSimulationFeeList []dtos.LoanSimulationFeeDTO

  for _, fee := range loan.TotalFees {
    loanSimulationFee := dtos.LoanSimulationFeeDTO{
      Number: fee.Number,
      Amounts: dtos.LoanSimulationFeeAmountsDTO{
        Initial: fee.Amount.Initial,
        Interest: fee.Amount.Interest,
        Payment: fee.Amount.Payment,
        Final: fee.Amount.Final,
      },
    }

    loanSimulationFeeList = append(loanSimulationFeeList, loanSimulationFee)
  }

  loanSimulationDto := dtos.LoanSimulationDTO{
    ID: uuid.New().String(),
    LoanAmount: loan.Amount,
    FeeAmount: loan.Fee,
    PeriodNumber: loan.Period,
    Fees: loanSimulationFeeList,
  }

  return loanSimulationDto
}
