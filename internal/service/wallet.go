package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/sumayu/testovoe2/internal/dto"
	"github.com/sumayu/testovoe2/internal/repository"
)

type WalletService struct {
	repo repository.WalletRepositoryInterface
}

func NewWalletService(repo repository.WalletRepositoryInterface) *WalletService {
	return &WalletService{repo: repo}
}

func (s *WalletService) ProcessTransaction(ctx context.Context, req dto.BalanceRequest) error {
	walletID, err := uuid.Parse(req.WalletID)
	if err != nil {
		return fmt.Errorf("invalid wallet ID: %v", err)
	}
	
	if req.OperationType != "DEPOSIT" && req.OperationType != "WITHDRAW" {
		return fmt.Errorf("invalid operation type")
	}
	
	amount := req.Amount
	if req.OperationType == "WITHDRAW" {
		amount = -amount
		balance, err := s.repo.GetBalance(ctx, walletID)
		if err != nil {
			return fmt.Errorf("failed to get balance: %v", err)
		}
		if balance < -amount {
			return fmt.Errorf("insufficient funds")
		}
	}
	
	if err := s.repo.UpdateBalance(ctx, walletID, amount); err != nil {
		return fmt.Errorf("update failed: %v", err)
	}
	return nil
}
func (s *WalletService) GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error) {
    return s.repo.GetBalance(ctx, walletID)
}