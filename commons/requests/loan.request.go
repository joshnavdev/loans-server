package request

type SimulateLoanRequest struct {
  Amount float64 `json:"amount" binding:"required"`
  AnnualPercentageRate float64 `json:"annualPercentageRate" binding:"required"`
  Period int `json:"period" binding:"required"`
}
