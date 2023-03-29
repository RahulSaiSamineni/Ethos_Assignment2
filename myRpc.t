MyRpc interface {
	GetBallance(username string) (availablebalance float64)
	Deposit(username string, amount float64) (depositamount float64)
	Withdraw(username string, amount1 float64) (withdrawnamount float64)
	Transfer(username string, username2 string, amount2 float64) (amount3 float64)
}
