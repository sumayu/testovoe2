package service

import (
	"context"
	"github.com/google/uuid"
	"github.com/sumayu/testovoe2/internal/dto"
	"github.com/stretchr/testify/mock"
)

type MockWalletService struct {
	mock.Mock
}

func (m *MockWalletService) ProcessTransaction(ctx context.Context, req dto.BalanceRequest) error {
	args := m.Called(ctx, req)
	return args.Error(0)
}

func (m *MockWalletService) GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error) {
	args := m.Called(ctx, walletID)
	return args.Get(0).(float64), args.Error(1)
}