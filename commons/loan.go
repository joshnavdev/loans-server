package common

import (
	"fmt"
	"math"

	request "github.com/joshnavdev/loans-server/commons/requests"
)

type LoanFeeAmount struct {
  Initial float64
  Fee float64
  Interest float64
  Payment float64
  Final float64
}

type LoanFee struct {
  ID string
  Number int
  Amount LoanFeeAmount
}

type Loan struct {
  Amount float64
  AnnualPercentageRate float64
  Period int
  MonthPercentageRate float64
  IA float64
  Fee float64
  TotalFees []LoanFee
}

func NewLoan(loan request.SimulateLoanRequest) Loan {
  return Loan{
    Amount: loan.Amount,
    AnnualPercentageRate: loan.AnnualPercentageRate,
    Period: loan.Period,
  }
}

func (loan *Loan) Simulate() {
  loan.CalculateMonthRate()
  loan.CalculateIA()
  loan.CalculateFee()
  loan.CalculateTotalFees()
}

func (loan *Loan) CalculateMonthRate() {
  var i float64 = 1 / 12.0
  loan.MonthPercentageRate = math.Pow(1 + loan.AnnualPercentageRate, i) - 1
}

func (loan *Loan) CalculateIA() {
  var sum float64 = 0

  // should use go parallel
  for i := 1; i <= loan.Period; i++ {
    iai := 1 / math.Pow(1 + loan.MonthPercentageRate, float64(i))

    sum += iai
  }

  loan.IA = sum
}

func (loan *Loan) CalculateFee() {
  loan.Fee = loan.Amount / loan.IA
}

func (loan *Loan) CalculateTotalFees() {
  var totalFees []LoanFee
  initialAmount := loan.Amount

  for i := 1; i <= loan.Period; i++ {
    fee := loan.Fee
    interest := initialAmount * loan.MonthPercentageRate
    payment := fee - interest
    finalAmount := initialAmount - payment

    loanFee := LoanFee{
      ID: fmt.Sprint(i),
      Number: i,
      Amount: LoanFeeAmount{
        Initial: initialAmount,
        Fee: fee,
        Interest: interest,
        Payment: payment,
        Final: finalAmount,
      },
    }

    initialAmount = finalAmount

    totalFees = append(totalFees, loanFee)
  }

  loan.TotalFees = totalFees
}
