package service

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/sumayu/testovoe2/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type mockWalletRepository struct {
	mock.Mock
}

func (m *mockWalletRepository) UpdateBalance(ctx context.Context, walletID uuid.UUID, amount float64) error {
	args := m.Called(ctx, walletID, amount)
	return args.Error(0)
}

func (m *mockWalletRepository) GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error) {
	args := m.Called(ctx, walletID)
	return args.Get(0).(float64), args.Error(1)
}

func TestWalletService_ProcessTransaction(t *testing.T) {
	testUUID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")

	mockRepo := new(mockWalletRepository)
	mockRepo.On("UpdateBalance", mock.Anything, testUUID, 100.0).Return(nil)

	service := NewWalletService(mockRepo)
	err := service.ProcessTransaction(context.Background(), dto.BalanceRequest{
		WalletID:      "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11",
		OperationType: "DEPOSIT",
		Amount:        100,
	})

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}