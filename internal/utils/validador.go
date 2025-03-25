package utils

import (
	z "github.com/Oudwins/zog"
)

var PersonAccountRequestValidador = z.Struct(z.Schema{
	"monthlyIncome": z.Float64().Required(),
	"Age":           z.Int32().Required(z.Message("Field must be required")),
	"FullName":      z.String().Max(265).Required(z.Message("Field must be required")),
	"Phone":         z.String().Required(z.Message("Field must be required")),
	"Email":         z.String().Email().Required(z.Message("Field must be required")),
	"Category":      z.String(),
	"Balance":       z.Float64().Required(z.Message("Field must be required")),
})

var PersonWithdrawRequestValidador = z.Struct(z.Schema{
	"withdrawValue": z.Float64().GT(0.0).Required(z.Message("This field must be greater than 0 and required")),
})

var PersonDepositRequest = z.Struct(z.Schema{
	"DepositValue": z.Float64().GT(0.0).Required(z.Message("This field must be greater than 0 and required")),
})
