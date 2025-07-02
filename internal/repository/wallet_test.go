
package repository

import (
	"context"
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestWalletRepository_GetBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("Ошибка создания мока БД: %v", err)
	}
	defer db.Close()
	repo := NewWalletRepository(db)
	testUUID := uuid.New()
	expectedBalance := 100.0

	t.Run("correct", func(t *testing.T) {
		rows := sqlmock.NewRows([]string{"balance"}).AddRow(expectedBalance)
		mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1").
			WithArgs(testUUID).
			WillReturnRows(rows)

		balance, err := repo.GetBalance(context.Background(), testUUID)

		assert.NoError(t, err)
		assert.Equal(t, expectedBalance, balance)
	})

	t.Run("no wallet", func(t *testing.T) {
		mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1").
			WithArgs(testUUID).
			WillReturnError(sql.ErrNoRows)

		_, err := repo.GetBalance(context.Background(), testUUID)
		assert.ErrorIs(t, err, sql.ErrNoRows)
	})

	assert.NoError(t, mock.ExpectationsWereMet())
}

func TestWalletRepository_UpdateBalance(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	repo := NewWalletRepository(db)
	testUUID := uuid.New()

	t.Run("update balance correct", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1 FOR UPDATE").
			WithArgs(testUUID).
			WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(50.0))
		mock.ExpectExec("UPDATE wallets SET balance = \\$1 WHERE id = \\$2").
			WithArgs(150.0, testUUID).
			WillReturnResult(sqlmock.NewResult(1, 1))
		mock.ExpectCommit()

		err := repo.UpdateBalance(context.Background(), testUUID, 100.0)
		assert.NoError(t, err)
	})

	t.Run("insufficient funds in wallet", func(t *testing.T) {
		mock.ExpectBegin()
		mock.ExpectQuery("SELECT balance FROM wallets WHERE id = \\$1 FOR UPDATE").
			WithArgs(testUUID).
			WillReturnRows(sqlmock.NewRows([]string{"balance"}).AddRow(50.0))
		mock.ExpectRollback()

		err := repo.UpdateBalance(context.Background(), testUUID, -100.0)
		assert.ErrorContains(t, err, "insufficient funds")
	})

	assert.NoError(t, mock.ExpectationsWereMet())
}