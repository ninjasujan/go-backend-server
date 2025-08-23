package model

type Auth struct {
	Id          string `json:"id"`
	Email       string `json:"email" validate:"required,email"`
	FirstName   string `json:"firstName" validate:"required"`
	Lastname    string `json:"LastName" validate:"required"`
	DateOfBirth string `json:"DOB"`
	Address     string `json:"address"`
}
