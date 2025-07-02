package repository
import (
	"context"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
)

type MockWalletRepository struct {
	mock.Mock
}

func (m *MockWalletRepository) UpdateBalance(ctx context.Context, walletID uuid.UUID, amount float64) error {
	args := m.Called(ctx, walletID, amount)
	return args.Error(0)
}

func (m *MockWalletRepository) GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error) {
	args := m.Called(ctx, walletID)
	return args.Get(0).(float64), args.Error(1)
}