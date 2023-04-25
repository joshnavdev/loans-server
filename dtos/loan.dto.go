package dtos

type LoanSimulationFeeAmountsDTO struct {
  Initial float64 `json:"initial"`
  Interest float64 `json:"interest"`
  Payment float64 `json:"payment"`
  Final float64 `json:"final"`
}

type LoanSimulationFeeDTO struct {
  Number int `json:"number"`
  Amounts LoanSimulationFeeAmountsDTO `json:"amounts"`
}

type LoanSimulationDTO struct {
  ID string `json:"id"`
  LoanAmount float64 `json:"loanAmount"`
  FeeAmount float64 `json:"feeAmount"`
  PeriodNumber int `json:"periodNumber"`
  Fees []LoanSimulationFeeDTO `json:"fees"`
}
