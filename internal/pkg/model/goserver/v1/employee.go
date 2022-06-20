package v1

import "github.com/go-playground/validator/v10"

type Employee struct {
	BaseModel

	// Required: true
	Username string `json:"username" gorm:"column:username;not null" binding:"required" validate:"min=1,max=32"`

	// Required: true
	WalletAddress string `json:"wallet_address" gorm:"colum:wallet_address;not null" binding:"required"`

	// Required: true
	Password string `json:"password,omitempty" gorm:"column:password;not null" binding:"required" validate:"min=5,max=128"`

	// Required: true
	Nickname string `json:"nickname" gorm:"column:nickname" binding:"required" validate:"required,min=1,max=30"`

	// Required: true
	Email string `json:"email" gorm:"column:email" binding:"required" validate:"required,email,min=1,max=100"`
}

// TableName maps to mysql table name.
func (u *Employee) TableName() string {
	return "employee"
}

// Validate the fields.
func (u *Employee) Validate() error {
	validate := validator.New()
	return validate.Struct(u)
}

/*
// EmployeeList is the whole list of all users which have been stored in stroage.
type EmployeeList struct {
	// May add TypeMeta in the future.
	// metav1.TypeMeta `json:",inline"`

	// Standard list metadata.
	// +optional
	metav1.ListMeta `json:",inline"`
	Items []*User `json:"items"`
}

*/
