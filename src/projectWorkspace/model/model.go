package model

import "time"

type User struct {
	UserName               string
	UserId                 string
	UserEMail              string
	AccountID              string
	PaymentDate            time.Time
	SubscriptionExpiryDate time.Time
	SubscriptionID         string
	PaymentMethod          string
}

type Account struct {
	AccountId   string
	AccountName string
	BillingID   string
	ServiceList []Services
}

//Services is a child schema of Account
type Services struct {
	ServiceName        string
	ServiceId          string
	ServiceMonthlyCost string //
	AnnualCost         string
}
