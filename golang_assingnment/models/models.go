package models

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_LIMIT int = 20000

func Msg() string {
	return "hello world"
}

type UserDetails struct {
	UserName  string
	AccountNo int
	PinNo     int
	Balance   int
}

func CheckBalance(obj UserDetails) {
	if obj.Balance > 500 {
		fmt.Println("Your Current Balance Is:", obj.Balance)
	} else if obj.Balance > 0 && obj.Balance < 500 {
		fmt.Println("Your Current Balance Is:", obj.Balance)
		fmt.Println("Warning: Please Maintain Balance of Minimum 500 Rs.")
	} else {
		fmt.Println("Your Current Balance Is:0")
	}
}

func DepositMoney(obj *UserDetails, ToBeDeposit string) {

	ToBeDeposit = strings.TrimSpace(ToBeDeposit)
	deposit, err := strconv.Atoi(ToBeDeposit)
	if err != nil {
		fmt.Println("error Occured:", err)
	} else {
		obj.Balance += deposit
		fmt.Println("Money Deposit Successful")
		fmt.Printf("Deposited Money:%d\n", deposit)
		fmt.Println("Balance After The Transaction:", obj.Balance)
	}

}

func WithdrawMoney(obj *UserDetails, ToBeWithdrawn string) {
	ToBeWithdrawn = strings.TrimSpace(ToBeWithdrawn)
	withdraw, err := strconv.Atoi(ToBeWithdrawn)

	if err != nil {
		fmt.Println("error Occured:", err)
	} else if withdraw < 0 {
		fmt.Println("Please enter valid amount")
	} else if obj.Balance < withdraw {
		fmt.Println("Cannot Perform the Transaction:Your Current Balance is Less than Required Amount")
	} else if withdraw > MAX_LIMIT {
		fmt.Println("Cannot Perform Transaction: Maximum Amount to Withdraw is:", MAX_LIMIT)
		return
	} else {
		obj.Balance -= withdraw
		fmt.Println("Balance After The Transaction:", obj.Balance)
	}
}
