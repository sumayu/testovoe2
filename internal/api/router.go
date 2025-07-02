	package api

	import (
		"github.com/gin-gonic/gin"
		"github.com/sumayu/testovoe2/internal/handler"
		"github.com/sumayu/testovoe2/internal/repository"
		"github.com/sumayu/testovoe2/internal/service"
		"database/sql"
	)

	func Router(db *sql.DB) *gin.Engine {
		r := gin.Default()
		walletRepo := repository.NewWalletRepository(db)
		walletService := service.NewWalletService(walletRepo)
		walletHandler := handler.NewWalletHandler(walletService)
		
		apiV1 := r.Group("/api/v1")
		{
			apiV1.POST("wallet", walletHandler.UpdateWalletBalance)
			apiV1.GET("wallets/:id", walletHandler.GetWalletBalance)
		}
		
		return r
	}