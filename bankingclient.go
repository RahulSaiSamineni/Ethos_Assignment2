package main
// import statements
import (
	"ethos/altEthos"
	"ethos/syscall"
	"ethos/myRpc"
	"log"
	"strconv"
)
// initializing the functions
func init(){
	myRpc.SetupMyRpcGetBallanceReply(getBallanceReply)
	myRpc.SetupMyRpcDepositReply(depositReply)
	myRpc.SetupMyRpcWithdrawReply(withdrawReply)
	myRpc.SetupMyRpcTransferReply(transferReply)
}
// fetching the balance from server and returning using the function below
func getBallanceReply(availablebalance float64) (myRpc.MyRpcProcedure) {
	log.Printf("Total Balance available in your account is: %f\n", availablebalance)
	return nil
}
// fetching the amount deposited from server and returning using the function below
func depositReply(depositamount float64) (myRpc.MyRpcProcedure) {
	log.Printf("Amount Deposited to your account is: %f\n", depositamount)
	return nil
}
// fetching the amount withdrawn from server and returning using the function below
func withdrawReply(withdrawnamount float64) (myRpc.MyRpcProcedure) {
	log.Printf("Amount Withdrawn from your account is: %f\n", withdrawnamount)
	return nil
}
// fetching the amount transferred from server and returning using the function below
func transferReply(amount3 float64) (myRpc.MyRpcProcedure) {
	log.Printf("The total amountTransfered is: "+ strconv.Itoa(int(amount3)))
	return nil
}

// main function
func main(){

	altEthos.LogToDirectory("test/bankingclient")
	var username = altEthos.GetUser() //fetching current logged in user using ethos
	var username2 = "bennett" // setting user2 for transfer
	log.Println("The current loggedin user name is: ",username)
	log.Println("Balance call")
	fd, status := altEthos.IpcRepeat("myRpc", "", nil)
	if status != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status)
		altEthos.Exit(status)
	}

	call := myRpc.MyRpcGetBallance{username} //call get ballance by passing user as parameter
	status = altEthos.ClientCall(fd, &call)
	if status != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status)
		altEthos.Exit(status)
	}

	log.Println("Deposit call")
	fd, status1 := altEthos.IpcRepeat("myRpc", "", nil)
	if status1 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	call1 := myRpc.MyRpcDeposit{username, 500} //calling deposit by passing user and amount to be deposited as parameters
	status1 = altEthos.ClientCall(fd, &call1)
	if status1 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status1)
		altEthos.Exit(status1)
	}

	log.Println("Withdraw call")
	fd, status2 := altEthos.IpcRepeat("myRpc", "", nil)
	if status2 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	call2 := myRpc.MyRpcWithdraw{username, 200} //calling withdraw by passing user and amount to be withdrawn as parameters
	status2 = altEthos.ClientCall(fd, &call2)
	if status2 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status2)
		altEthos.Exit(status2)
	}

	log.Println("Transfer call")
	fd, status3 := altEthos.IpcRepeat("myRpc", "", nil)
	if status3 != syscall.StatusOk {
		log.Printf("Ipc failed: %v\n", status3)
		altEthos.Exit(status3)
	}

	call3 := myRpc.MyRpcTransfer{username,username2, 200} //calling transfer by passing 2 usernames and amount to be transferred as parameters
	status3 = altEthos.ClientCall(fd, &call3)
	if status3 != syscall.StatusOk {
		log.Printf("clientCall failed: %v\n", status3)
		altEthos.Exit(status3)
	}


	log.Println("bankingclient: done")
}