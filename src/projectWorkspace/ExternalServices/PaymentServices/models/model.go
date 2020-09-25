package models

import (
	"time"
)

type UserPayment struct {
	ReferenceNumber int `json:"reference number"`
	PaymentDate time.Time
	PaymentMode string `json:"mode"`
	Amount float32  `json:"amount"`
	Status string `json:"Status:"`
}

type Response struct {
	Message string `json:"message"`
}