package v1

import (
	"context"
	"fmt"
	"github.com/FuradWho/Cherry/internal/apiserver/store"
	"github.com/FuradWho/Cherry/internal/pkg/errno"
	metav1 "github.com/FuradWho/Cherry/internal/pkg/model/goserver/v1"
	"regexp"
)

type EmployeeSrv interface {

	// , user *v1.User, opts metav1.CreateOptions

	Register(ctx context.Context, employee *metav1.Employee) error
	Get(ctx context.Context, walletAddress string) (*metav1.Employee, error)
	NFTApplication(ctx context.Context, walletAddress, password string) (string, error)
	NFTAudit(ctx context.Context, nftAddress string) (string, error)
}

type employeeService struct {
	store store.Factory
}

func newEmployees(srv *service) *employeeService {
	return &employeeService{store: srv.store}
}

func (s *employeeService) Register(ctx context.Context, employee *metav1.Employee) error {
	err := s.store.Employees().Register(ctx, employee)
	if err != nil {
		if match, _ := regexp.MatchString("Duplicate entry '.*' for key 'username'", err.Error()); match {
			return errno.ErrUserAlreadyExist
		}
		return err
	}
	return nil
}

func (s *employeeService) Get(ctx context.Context, walletAddress string) (*metav1.Employee, error) {
	employee, err := s.store.Employees().Get(ctx, walletAddress)
	if err != nil {
		return nil, err
	}
	return employee, nil
}

func (s *employeeService) NFTApplication(ctx context.Context, walletAddress, password string) (string, error) {

	employee, err := s.Get(ctx, walletAddress)
	if err != nil {
		return "", err
	}

	fmt.Println(employee.WalletAddress)

	if employee == nil || employee.Password != password {
		return "", err
	}

	//TODO test nft address is 11111111111111111111111111

	nftAddr := "11111111111111111111111111"

	err = s.store.Employees().SaveNFT(ctx, &metav1.NFT{
		WalletAddress: walletAddress,
		NFTAddress:    nftAddr,
	})
	if err != nil {
		return "", err
	}

	return nftAddr, nil

}

func (s *employeeService) NFTAudit(ctx context.Context, nftAddress string) (string, error) {
	nft, err := s.store.Employees().GetNFT(ctx, nftAddress)
	if err != nil || nft == nil {
		return "-1", err
	}

	//TODO get nft record on chain

	return "1", err

}
