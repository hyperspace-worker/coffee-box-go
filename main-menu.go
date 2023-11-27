package main

import (
	"bufio"
	"fmt"
	"os"
)

func callMainMenu(glasses *int, userBalance *float32, cashBalance *float32, availablePinInputAttempts int) {

	r := bufio.NewReader(os.Stdin)

	for true {
		choiseOption := 0

		showMainMenu(*userBalance, *glasses)

		fmt.Print("Please, choise option: ")

		res, err := selectOption(r)

		fmt.Printf("Selected option is %v, err is %v\n", res, err)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			clearScreen()
			continue
		}

		choiseOption = res

		clearScreen()

		switch choiseOption {
		case 1:
			giveCoffeeToUser(userBalance, PRICE_CAPPUCCINO, glasses)
		case 2:
			giveCoffeeToUser(userBalance, PRICE_ESPRESSO, glasses)
		case 3:
			giveCoffeeToUser(userBalance, PRICE_LATTE, glasses)
		case 4:
			callCashDepositMenu(userBalance, cashBalance)
		case 5:
			switch checkAccess(availablePinInputAttempts) {
			case 0:
				showMessage("You cancelled operation")
			case 1:
				callServiceMenu(glasses, cashBalance, userBalance, availablePinInputAttempts)
			case -1:
				fmt.Println()
				showSymbolsRow()
				showSymbolsRowWithMessage("The coffe machine is blocked!", ROW_LENGTH)
				showSymbolsRowWithMessage("Reason: too many PIN input attempts.", ROW_LENGTH)
				showSymbolsRowWithMessage("Please, call our contact manager to unlock:", ROW_LENGTH)
				showSymbolsRowWithMessage(MANAGER_CONTACTS, ROW_LENGTH)
				showSymbolsRow()
				return
			}
		default:
			showWrongInputMessage()
		}
	}
}

func showMainMenu(userBalance float32, glasses int) {
	showHeader("ESPRESSO BIANCCI")
	showSymbolsRowWithMessage("MAIN MENU", ROW_LENGTH)
	showSymbolsRow()
	fmt.Printf("%-25v %v BYN\n", "Cash balance:", userBalance)
	showSymbolsRow()
	fmt.Printf("%-25v %v\n", "Number of glasses:", glasses)
	showSymbolsRow()

	showSymbolsRowWithMessage("Select coffee", ROW_LENGTH)
	showSymbolsRow()
	fmt.Printf("*%-22v %-5v %-12v\n", "1. Cappuccino", PRICE_CAPPUCCINO, "*")
	fmt.Printf("*%-20v %-7v %-12v\n", "2. Espresso", PRICE_ESPRESSO, "*")
	fmt.Printf("*%-17v %-10v %-12v\n", "3. Latte", PRICE_LATTE, "*")
	showSymbolsRow()

	fmt.Printf("*%-24v %-15v\n", "4. Cash deposit", "*")
	showSymbolsRow()
	fmt.Printf("*%-19v %-20v\n", "5. Service", "*")
	showSymbolsRow()
}

func showHeader(headerText string) {
	showSymbolsRow()
	showSymbolsRow()
	fmt.Println()
	showSymbolsRowWithMessage(headerText, ROW_LENGTH)
	fmt.Println()
	showSymbolsRow()
	showSymbolsRow()
}

func showCashDepositMenu() {
	showSymbolsRow()
	showSymbolsRowWithMessage("CASH DEPOSIT (BYN)", ROW_LENGTH)
	showSymbolsRow()
	fmt.Printf("* 0. Exit %-25v\n", "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_05, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_1, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_2, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_5, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_10, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_20, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_50, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_100, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_200, "*")
	fmt.Printf("*%-15v BYN %-25v\n", BYN_BILL_500, "*")
	showSymbolsRow()
}

func showCoffeeIsPurchased() {
	showMessage("Take your coffee, please!")
}

func showNotEnoughMoneyWarning() {
	showMessage("Sorry, you don't have enough money!\nYou need to put some cash\nin coffe machine")
}

func showNoGlassesWarning() {
	showMessage(fmt.Sprintf("Sorry we don't have glasses! Please, call our manager: %v", MANAGER_CONTACTS))
}

func showMoneyFromUser(byn float32) {
	showMessage(fmt.Sprintf(" You put in coffe machine %0.1f BYN \n", byn))
}

func showSymbolsRow() {
	for i := 0; i < ROW_LENGTH; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
}

func showSymbolsRowWithMessage(message string, rowLength int) {
	fmt.Print("* ")

	freeSpace := (rowLength - 4 - len(message))
	halfOfFreeSpace := freeSpace / 2

	for i := 0; i < halfOfFreeSpace; i++ {
		fmt.Print(" ")
	}

	fmt.Print(message)

	for i := 0; i < halfOfFreeSpace; i++ {
		fmt.Print(" ")
	}

	if freeSpace%2 == 1 {
		fmt.Print(" ")
	}

	fmt.Printf(" *\n")
}

func callCashDepositMenu(userBalance *float32, cashBalance *float32) int {

	r := bufio.NewReader(os.Stdin)

	for {
		choiseOption := -1

		showCashDepositMenu()

		fmt.Println()
		fmt.Println("Please, select how much money do you")
		fmt.Println("want to put in coffee machine: ")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			continue
		}

		choiseOption = res

		clearScreen()

		switch choiseOption {
		case 0:
			return 0
		case 1:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_05)
		case 2:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_1)
		case 3:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_2)
		case 4:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_5)
		case 5:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_10)
		case 6:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_20)
		case 7:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_50)
		case 8:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_100)
		case 9:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_200)
		case 10:
			getMoneyFromUser(userBalance, cashBalance, BYN_BILL_500)
		default:
			showWrongInputMessage()
		}
	}
}
