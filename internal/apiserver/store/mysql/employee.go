package mysql

import (
	"context"
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
