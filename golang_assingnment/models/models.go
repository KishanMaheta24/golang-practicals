package models

import (
	"fmt"
	"strconv"
	"strings"
)

const MAX_LIMIT int = 20000
const MAX_LIMIT_DEP int = 250000

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
	if obj.Balance >= 500 {
		fmt.Println("Your Current Balance Is:", obj.Balance)
	} else if obj.Balance > 0 && obj.Balance < 500 {
		fmt.Println("Your Current Balance Is:", obj.Balance)
		fmt.Println("Warning: Please Maintain Balance of Minimum 500 Rs.")
	} else {
		fmt.Println("Your Current Balance Is:", obj.Balance)
	}
}

func DepositMoney(obj *UserDetails, ToBeDeposit string) {

	ToBeDeposit = strings.TrimSpace(ToBeDeposit)
	_, err1 := strconv.ParseFloat(ToBeDeposit, 5)
	//fmt.Println(reflect.TypeOf(isFloat))
	deposit, err := strconv.Atoi(ToBeDeposit)
	if err != nil && err1 == nil {
		fmt.Println("Cannot enter amount in Float Values,Please Enter Whole values only!")
	} else if err != nil {
		fmt.Println("error Occured: Please enter Valid Amount")
	} else if deposit < 0 {
		fmt.Println("Please enter valid amount")
	} else if (deposit % 500) != 0 {
		fmt.Println("Please Enter Money in Multiple of [500] ")
	} else if deposit >= MAX_LIMIT_DEP {
		fmt.Println("Sorry, You Cannot Deposit More Than ", MAX_LIMIT_DEP, " in Single Transaction")
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
	_, err1 := strconv.ParseFloat(ToBeWithdrawn, 5)
	if err != nil && err1 == nil {
		fmt.Println("Cannot enter amount in Float Values,Please Enter Whole values only!")
	} else if err != nil {
		fmt.Println("error Occured:", "Please enter Valid Amount")
	} else if withdraw < 0 {
		fmt.Println("Please enter valid amount")
	} else if withdraw%500 != 0 {
		fmt.Println("Please Enter Money in Multiple of [500]")
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
