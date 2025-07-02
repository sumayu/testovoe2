package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/google/uuid"
)

type WalletRepositoryInterface interface {
	UpdateBalance(ctx context.Context, walletID uuid.UUID, amount float64) error
	GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error)
}

type WalletRepository struct {
	db *sql.DB
}

func NewWalletRepository(db *sql.DB) *WalletRepository {
	return &WalletRepository{db: db}
}	

func (r *WalletRepository) UpdateBalance(ctx context.Context, walletID uuid.UUID, amount float64) error {
    tx, err := r.db.BeginTx(ctx, nil)
    if err != nil {
        return err
    }
    defer tx.Rollback()

    var currentBalance float64
    err = tx.QueryRowContext(ctx,
        "SELECT balance FROM wallets WHERE id = $1 FOR UPDATE",
        walletID,
    ).Scan(&currentBalance)
    if err != nil {
        return err
    }

    newBalance := currentBalance + amount
    if newBalance < 0 {
        return errors.New("insufficient funds")
    }

    _, err = tx.ExecContext(ctx,
        "UPDATE wallets SET balance = $1 WHERE id = $2",
        newBalance, walletID,
    )
    if err != nil {
        return err
    }

    return tx.Commit()
}

func (r *WalletRepository) GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error) {
	var balance float64
	err := r.db.QueryRowContext(ctx,
		"SELECT balance FROM wallets WHERE id = $1",
		walletID,
	).Scan(&balance)

	return balance, err
}