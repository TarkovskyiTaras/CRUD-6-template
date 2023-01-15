package domain

import "time"

type PersonInfo struct {
	ID          int       `json:"id"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	DOB         time.Time `json:"dob"`
	HomeAddress string    `json:"home_address"`
	CellPhone   string    `json:"cellphone"`
}

type UpdatePersonInput struct {
	FirstName   *string    `json:"first_name"`
	LastName    *string    `json:"last_name"`
	DOB         *time.Time `json:"dob"`
	HomeAddress *string    `json:"home_address"`
	CellPhone   *string    `json:"cellphone"`
}
