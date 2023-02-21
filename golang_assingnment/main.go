package main

import (
	"bufio"
	"fmt"
	"models/models/models"
	"os"
	"strconv"
	"strings"
)

type UserDetails = models.UserDetails

var TotalAccounts = []UserDetails{
	{"kishan", 1234567890, 1234, 1000},
	{"dummy1", 9876543210, 4321, 5000},
}

func handlePanicMode() {
	if a := recover(); a != nil {

		fmt.Println("RECOVER", a)
	}
}

func main() {
	scanner := bufio.NewReader(os.Stdin)
	defer handlePanicMode()
	fmt.Println()

	//models.Msg()
enterCredentials:
	fmt.Println("============")
	fmt.Println("UserName:kishan, Account No:1234567890, Pin:1234, Balance:1000")
	fmt.Println("UserName:dummy1, Account No:9876543210, Pin:4321, Balance:5000")
	fmt.Println("============")
	fmt.Println("Welcome To Simform ATM ")

	fmt.Println("Please enter your ac no:")

	accountNo, err1 := scanner.ReadString('\n')
	accountNo = strings.TrimSpace(accountNo)

	fmt.Println("enter your pin:")
	userPin, err2 := scanner.ReadString('\n')
	userPin = strings.TrimSpace(userPin)

	if checkError(err1) || checkError(err2) {
		fmt.Println("Internal error: Try Again")
		goto enterCredentials
	}

	credentials, status := checkCredentials(accountNo, userPin)

	//fmt.Println(credentials, status)

	if status == "Failed" {
		fmt.Println("Invalid Account No. or Pin.... Try Again")

		//fmt.Println()
		goto enterCredentials
	} else if status == "WP" {

		fmt.Println("Wrong Password")
		goto enterCredentials
	} else if status == "NAP" {
		//fmt.Println("panicking")
		fmt.Println("No Such Account Found")
		goto enterCredentials
	}

	for {
	Choice:
		fmt.Println()
		fmt.Printf("Welcome,%s...!\n", credentials.UserName)
		fmt.Println("===================")
		fmt.Println("Enter 1 to Check Balance")
		fmt.Println("Enter 2 to Deposit")
		fmt.Println("Enter 3 to Withdraw")
		fmt.Println("Type E/e to exit")
		fmt.Print("Your choice:")
		userChoice, err := scanner.ReadString('\n')

		if err != nil {
			fmt.Println("ar error occurred while execution:", err)
			goto Choice
		} else {
			userChoice = strings.TrimSpace(userChoice)
		}

		switch userChoice {

		case "1":
			models.CheckBalance(credentials)
		case "2":
			fmt.Print("Enter Money To Deposit:")
			deposit, err := scanner.ReadString('\n')
			if checkError(err) {
				fmt.Println("Internal error: Try Again")
			} else {
				models.DepositMoney(&credentials, deposit)
			}

		case "3":
			fmt.Print("Enter Money To Withdraw:")
			withdraw, err := scanner.ReadString('\n')
			if err != nil {
				fmt.Println("Error Occurred.. Try Again")
			} else {
				models.WithdrawMoney(&credentials, withdraw)
			}
		case "E":
			fmt.Println("Thank you... Have a good day")
			os.Exit(1)
		case "e":
			fmt.Println("Thank you... Have a good day")
			os.Exit(1)
		default:
			fmt.Println("asd")
		}

	}
}

func checkError(err error) bool {
	if err != nil {
		return true
	} else {
		return false
	}
}

func checkCredentials(accNo string, pin string) (UserDetails, string) {
	//WP-wrong password
	//NAP - no account found

	accNo1, _ := strconv.Atoi(accNo)
	pin1, _ := strconv.Atoi(pin)

	flag1 := false
	//flag2 := false
	for _, obj := range TotalAccounts {
		//fmt.Println(obj.AccountNo)
		if obj.AccountNo == accNo1 {
			flag1 = true
			if obj.PinNo != pin1 {
				return UserDetails{}, "WP"
			} else {
				return obj, "Success"
			}
		}

	}
	if !flag1 {
		return UserDetails{}, "NAP"
	} else {
		return UserDetails{}, "Failed"
	}
}
