package mysql

import (
	"context"
	"errors"
	"github.com/FuradWho/Cherry/internal/pkg/errno"
	metav1 "github.com/FuradWho/Cherry/internal/pkg/model/goserver/v1"
	"gorm.io/gorm"
)

type employees struct {
	db *gorm.DB
}

func newEmployees(ds *datastore) *employees {
	return &employees{ds.db}
}

// Register register a new user account.
func (u *employees) Register(ctx context.Context, employee *metav1.Employee) error {
	return u.db.Create(&employee).Error
}

// Get return an employee by the employee wallet_address.
func (u *employees) Get(ctx context.Context, walletAddress string) (*metav1.Employee, error) {
	employee := &metav1.Employee{}

	err := u.db.Where("wallet_address = ?", walletAddress).First(employee).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}

		return nil, err
	}

	return employee, nil

}

func (u *employees) SaveNFT(ctx context.Context, nft *metav1.NFT) error {
	return u.db.Create(&nft).Error
}

func (u *employees) GetNFT(ctx context.Context, nftAddress string) (*metav1.NFT, error) {
	nft := &metav1.NFT{}

	err := u.db.Where("nft_address = ?", nftAddress).First(nft).Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errno.ErrUserNotFound
		}

		return nil, err
	}

	return nft, nil
}
