package main

//import statements
import (
	"ethos/syscall"
	"ethos/altEthos"
	"ethos/myRpc"
	"log"
)
//created a map with users and their account balance
var users = map[string]float64{"me":2700, "bennett":1000, "jlong":1500, "gabriel":2900,}

// initializing the functions
func init() {
	myRpc.SetupMyRpcGetBallance(getBallance)
	myRpc.SetupMyRpcDeposit(deposit)
	myRpc.SetupMyRpcWithdraw(withdraw)
	myRpc.SetupMyRpcTransfer(transfer)
}

//function for fetching the account balance based on username.
func getBallance(username string)(myRpc.MyRpcProcedure) {
	log.Println("The total balance available in your account is:", users[username])
	return &myRpc.MyRpcGetBallanceReply{users[username]}  //returning the balance for user to client
}
//function to deposit money into user account by passing username and amount to be deposited as parameters.
func deposit(username string, amount float64) (myRpc.MyRpcProcedure) {
	log.Println("Amount Deposited Successfully: ",amount)
	users[username] = users[username] + amount //updating the account balance of user after deposit.
	log.Println("Total balance available after deposit: ",users[username])
	return &myRpc.MyRpcDepositReply{amount} //returning the amount that is deposited.
}
//function to withdraw money from user account by passing username and amount to be withdrawn as parameters.
func withdraw(username string, amount1 float64) (myRpc.MyRpcProcedure){
	//Returning insufficient funds if account balance is less than requested amount
	if(amount1 >= users[username]){
		log.Println("Insufficient funds")
		return &myRpc.MyRpcWithdrawReply{0}
	}
	users[username] = users[username] - amount1 //else updating the account balance byreducting the withdrawn amount.
	log.Println("The amount withdrawn from your account is: ", amount1) 
	log.Println("The total balance available after withdrawn is: ", users[username])
	return &myRpc.MyRpcWithdrawReply{amount1} //returning withdrawn amount.
}
//function to transfer money from one user account to other user account by passing 2 user parameters and transfered amount parameter as parameters.
func transfer(username string, username2 string, amount2 float64) (myRpc.MyRpcProcedure){
	log.Println("The total amount transferred is:", amount2)
	users[username] = users[username] - amount2 //updating the balance of user who transferred the amount.
	users[username2] = users[username2] + amount2 //updating the balance of user who recieved the amount.
	log.Println("The total amount available after amount sent is: ", users[username])
	return &myRpc.MyRpcTransferReply{amount2} //returning the amount that is transferred.
}
// Main function
func main() {
	altEthos.LogToDirectory("test/bankingserver")

	listeningFd, status := altEthos.Advertise("myRpc")
	if status != syscall.StatusOk {
		log.Println("Advertising service failed: ", status)
		altEthos.Exit(status)
	}

	for {
		_, fd, status := altEthos.Import(listeningFd)
		if status != syscall.StatusOk {
			log.Printf("Error calling Import: %v\n", status)
			altEthos.Exit(status)
		}

		log.Println("new connection accepted")

		t := myRpc.MyRpc{}
		altEthos.Handle(fd, &t)
	}
}