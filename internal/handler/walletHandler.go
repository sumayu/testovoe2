package handler

import (
	"context"
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sumayu/testovoe2/internal/dto"
)

type WalletHandler struct {
	walletService WalletService
}

type WalletService interface {
	ProcessTransaction(ctx context.Context, req dto.BalanceRequest) error
	GetBalance(ctx context.Context, walletID uuid.UUID) (float64, error)
}

func NewWalletHandler(walletService WalletService) *WalletHandler {
	return &WalletHandler{walletService: walletService}
}

func (h *WalletHandler) UpdateWalletBalance(c *gin.Context) {
	var req dto.BalanceRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Incorrect request format"})
		return
	}

	if req.OperationType != "DEPOSIT" && req.OperationType != "WITHDRAW" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Incorrect operation, available operations: DEPOSIT and WITHDRAW",
		})
		return
	}

	if err := h.walletService.ProcessTransaction(c.Request.Context(), req); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"status": "success"})
}

func (h *WalletHandler) GetWalletBalance(c *gin.Context) {
	walletUUIDstr := c.Param("id")
	walletUUID, err := uuid.Parse(walletUUIDstr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid wallet ID format"})
		return
	}

	balance, err := h.walletService.GetBalance(c.Request.Context(), walletUUID)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "Wallet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get balance"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"wallet_id": walletUUID,
		"balance":   balance,
	})
}