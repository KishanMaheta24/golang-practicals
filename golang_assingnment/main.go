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

var count_pass int = 0
var count_acno int = 0
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
	if count_acno == 3 {
		fmt.Println("Maximum number of attempts reached while entering Account Number... Try again later")
		os.Exit(1)
	}
	if count_pass == 3 {
		fmt.Println("Maximum number of attempts reached while entering PIN... Try again later")
		os.Exit(1)
	}

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
		fmt.Println("Invalid Error Occurred.... Try Again later")
		os.Exit(1)
	} else if status == "WP" {
		count_pass++
		fmt.Println("Wrong Password, Attempts Remaining:", 3-count_pass)
		goto enterCredentials
	} else if status == "NAP" {
		count_acno++
		fmt.Println("No Such Account Found,Attempts Remaining:", 3-count_acno)
		goto enterCredentials
	} else if status == "VAC" {
		count_acno++
		fmt.Println("Please Enter Valid Account Number [in Number Only],Attempts Remaining:", 3-count_acno)
		goto enterCredentials
	} else if status == "VP" {
		count_pass++
		fmt.Println("Please Enter Valid Pin [in Number Only], Attempts Remaining:", 3-count_pass)
		goto enterCredentials
	} else if status == "NVL" {
		count_acno++
		fmt.Println("Please Enter Account number in Length of 10,Attempts Remaining:", 3-count_acno)
		goto enterCredentials
	} else if status == "NVLP" {
		count_pass++
		fmt.Println("Please Enter PIN number in Length of 4, Attempts Remaining:", 3-count_pass)
		goto enterCredentials
	} else {
		fmt.Printf("Welcome,%s...!\n", credentials.UserName)
	}

	for {

		fmt.Println()

	Choice:
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
			fmt.Println("Thank You, Please Remove Your Card")
			os.Exit(1)
		case "e":
			fmt.Println("Thank You, Please Remove Your Card")
			os.Exit(1)
		default:
			fmt.Println("Incorrect Choice Entered, Please Enter a Valid Choice")
		}
	ContinueChoice:
		fmt.Println("\nDo You Want To Continue:(Y/N):")
		cont, err := scanner.ReadString('\n')
		cont = strings.TrimSpace(cont)
		if err != nil {
			fmt.Println("error occured Please Try Again later")
		} else {
			if cont == "Y" || cont == "y" {
				continue
			} else if cont == "n" || cont == "N" {
				fmt.Println("Thank You, Please Remove Your Card")
				os.Exit(1)
			} else {
				fmt.Println("Enter Valid Choice")
				goto ContinueChoice
			}
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

	accNo1, err := strconv.Atoi(accNo)
	pin1, err1 := strconv.Atoi(pin)

	if err != nil {
		return UserDetails{}, "VAC" //valid account number
	}

	if err1 != nil {
		return UserDetails{}, "VP" //valid pin
	}

	if len(accNo) != 10 {
		return UserDetails{}, "NVL" //not a valid length
	}

	if len(pin) != 4 {
		return UserDetails{}, "NVLP" //valid length pin
	}
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
