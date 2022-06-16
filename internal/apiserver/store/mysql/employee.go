package mysql

import (
	"context"
	"gorm.io/gorm"
)

type employees struct {
	db *gorm.DB
}

func newUsers(ds *datastore) *employees {
	return &employees{ds.db}
}

// Create creates a new user account.
func (u *employees) Create(ctx context.Context) error {
	// u.db.Create(&employee).Error
	return nil
}
