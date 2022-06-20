package v1

import "github.com/go-playground/validator/v10"

type NFT struct {
	BaseModel

	// Required: true
	WalletAddress string `json:"wallet_address" gorm:"colum:wallet_address;not null" binding:"required"`
	NFTAddress    string `json:"nft_address" gorm:"colum:nft_address;not null" binding:"required"`
}

// TableName maps to mysql table name.
func (n *NFT) TableName() string {
	return "nft"
}

// Validate the fields.
func (n *NFT) Validate() error {
	validate := validator.New()
	return validate.Struct(n)
}
