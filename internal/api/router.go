package api

import (
	"github.com/gin-gonic/gin"
)

func Router() *gin.Engine {
	//сделать группы обработчиков api и v1
// TO-DO сделать обработчик POST api/v1/wallet с json { будет запускаться функция handler.UpdateWalletBalance
//valletId: UUID,
//operationType: DEPOSIT or WITHDRAW,
//amount: 1000
//} 
// сделать обработчик GET api/v1/wallets/{WALLET_UUID} ( будет выводить фукнцию из пакета handler.GetWalletBalance)

}