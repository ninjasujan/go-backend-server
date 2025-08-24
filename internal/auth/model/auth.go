package model

type Auth struct {
	Id          string `json:"id" gorm:"id"`
	Email       string `json:"email" validate:"required,email" gorm:"email"`
	FirstName   string `json:"firstName" validate:"required" gorm:"first_name"`
	LastName    string `json:"lastName" validate:"required" gorm:"last_name"`
	DateOfBirth string `json:"DOB" gorm:"dob"`
	Address     string `json:"address" gorm:"address"`
}
