package api

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/sumayu/testovoe2/internal/bd"
)

func TestRouter(t *testing.T) {
	db,_:= bd.Database()
	if db == nil {
		t.Fatal("Failed to connect to database")
	}
	defer db.Close()
	router := Router(db)
	tests := []struct {
		name           string
		method         string
		path           string
		expectedStatus int
	}{
		{
			name:           "POST wallet endpoint exists",
			method:         "POST",
			path:           "/api/v1/wallet",
			expectedStatus: http.StatusBadRequest, 
		},
		{
			name:           "GET wallet balance endpoint exists",
			method:         "GET",
			path:           "/api/v1/wallets/123",
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := httptest.NewRecorder()
			req, _ := http.NewRequest(tt.method, tt.path, nil)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)
		})
	}
}