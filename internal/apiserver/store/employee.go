package store

import (
	"context"

	metav1 "github.com/FuradWho/Cherry/internal/pkg/model/goserver/v1"
)

type EmployeeStore interface {
	Register(ctx context.Context, employee *metav1.Employee) error
	Get(ctx context.Context, walletAddress string) (*metav1.Employee, error)
	SaveNFT(ctx context.Context, nft *metav1.NFT) error
	GetNFT(ctx context.Context, nftAddress string) (*metav1.NFT, error)
	// NFTApplication(ctx context.Context) error
}
