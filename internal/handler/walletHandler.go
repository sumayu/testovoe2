package handler

import "github.com/gin-gonic/gin"

//valletId: UUID
//operationType: DEPOSIT or WITHDRAW,
//amount: 1000

func UpdateWalletBalance(c *gin.Context)  {
	var req struct {
		WalletID      string  `json:"walletId"`       
        OperationType string  `json:"operationType"` 
        Amount        float64 `json:"amount"`         
	}
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(400, "Incorrect request format")
	}
	  if req.OperationType != "DEPOSIT" && req.OperationType != "WITHDRAW" {
        c.JSON(400, gin.H{"error": "Incorrect operation, available operation: DEPOSIT and WITHDRAW"})
        return
    }
		//to-do сделать в db функцию которая сможет принимать в себя структуру req, сравнить ее данные и изменить сумму
}

func GetWalletBalance()  {
	
}