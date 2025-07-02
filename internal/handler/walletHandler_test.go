package handler

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/sumayu/testovoe2/internal/dto"
	"github.com/stretchr/testify/assert"
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

func TestWalletHandler_UpdateWalletBalance(t *testing.T) {
	gin.SetMode(gin.TestMode)

	tests := []struct {
		name           string
		requestBody    string
		mockSetup      func(*MockWalletService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:        "successful deposit",
			requestBody: `{"walletId": "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11", "operationType": "DEPOSIT", "amount": 100.5}`,
			mockSetup: func(m *MockWalletService) {
				m.On("ProcessTransaction", mock.Anything, mock.MatchedBy(func(req dto.BalanceRequest) bool {
					return req.WalletID == "a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11" && req.OperationType == "DEPOSIT" && req.Amount == 100.5
				})).Return(nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"success"}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockWalletService)
			tt.mockSetup(mockService)
			handler := NewWalletHandler(mockService)
			router := gin.New()
			router.POST("/wallet", handler.UpdateWalletBalance)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("POST", "/wallet", strings.NewReader(tt.requestBody))
			req.Header.Set("Content-Type", "application/json")
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}

func TestWalletHandler_GetWalletBalance(t *testing.T) {
	gin.SetMode(gin.TestMode)
	testUUID := uuid.MustParse("a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11")

	tests := []struct {
		name           string
		walletID       string
		mockSetup      func(*MockWalletService)
		expectedStatus int
		expectedBody   string
	}{
		{
			name:     "successful get balance",
			walletID: testUUID.String(),
			mockSetup: func(m *MockWalletService) {
				m.On("GetBalance", mock.Anything, testUUID).Return(100.5, nil)
			},
			expectedStatus: http.StatusOK,
			expectedBody:   `{"wallet_id":"a0eebc99-9c0b-4ef8-bb6d-6bb9bd380a11","balance":100.5}`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockService := new(MockWalletService)
			tt.mockSetup(mockService)
			handler := NewWalletHandler(mockService)
			router := gin.New()
			router.GET("/wallet/:id", handler.GetWalletBalance)
			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/wallet/"+tt.walletID, nil)
			router.ServeHTTP(w, req)
			assert.Equal(t, tt.expectedStatus, w.Code)
			assert.JSONEq(t, tt.expectedBody, w.Body.String())
			mockService.AssertExpectations(t)
		})
	}
}