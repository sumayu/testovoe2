package service

import (
	"fmt"
	"github.com/google/uuid"
	"github.com/sumayu/testovoe2/internal/dto"
	"github.com/sumayu/testovoe2/internal/repository"
)

func ProcessTransaction(req dto.BalanceRequest) error {
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
		balance, err := repository.GetBalance(walletID)
		if err != nil {
			return fmt.Errorf("failed to get balance: %v", err)
		}
		if balance < -amount {
			return fmt.Errorf("insufficient funds")
		}
	}
	if err := repository.UpdateBalance(walletID, amount); err != nil {
		return fmt.Errorf("update failed: %v", err)
	}
	return nil
}