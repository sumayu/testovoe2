package dto
type BalanceRequest struct {
    WalletID      string  `json:"walletId"`
    OperationType string  `json:"operationType"` 
    Amount        float64 `json:"amount"`
}