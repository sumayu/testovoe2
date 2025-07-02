package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
)

type MockWalletHandler struct {
	mock.Mock
}

func (m *MockWalletHandler) UpdateWalletBalance(c *gin.Context) {
	m.Called(c)
}

func (m *MockWalletHandler) GetWalletBalance(c *gin.Context) {
	m.Called(c)
}