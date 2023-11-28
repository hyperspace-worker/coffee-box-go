package main

import (
	"bufio"
	"fmt"
	"os"
)

func callMainMenu(glasses *int, w *Wallet, availablePinInputAttempts int) {

	r := bufio.NewReader(os.Stdin)

	for true {
		choiseOption := 0

		showMainMenu(w, *glasses)

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
			giveCoffeeToUser(w, PRICE_CAPPUCCINO, glasses)
		case 2:
			giveCoffeeToUser(w, PRICE_ESPRESSO, glasses)
		case 3:
			giveCoffeeToUser(w, PRICE_LATTE, glasses)
		case 4:
			callCashDepositMenu(w)
		case 5:
			switch checkAccess(availablePinInputAttempts) {
			case 0:
				showMessage("You cancelled operation")
			case 1:
				callServiceMenu(glasses, w, availablePinInputAttempts)
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

func showMainMenu(w *Wallet, glasses int) {
	showHeader("ESPRESSO BIANCCI")
	showSymbolsRowWithMessage("MAIN MENU", ROW_LENGTH)
	showSymbolsRow()
	fmt.Printf("%-25v %v BYN\n", "Cash balance:", w.balance)
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

func callCashDepositMenu(w *Wallet) int {

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
			depositMoney(w, BYN_BILL_05)
		case 2:
			depositMoney(w, BYN_BILL_1)
		case 3:
			depositMoney(w, BYN_BILL_2)
		case 4:
			depositMoney(w, BYN_BILL_5)
		case 5:
			depositMoney(w, BYN_BILL_10)
		case 6:
			depositMoney(w, BYN_BILL_20)
		case 7:
			depositMoney(w, BYN_BILL_50)
		case 8:
			depositMoney(w, BYN_BILL_100)
		case 9:
			depositMoney(w, BYN_BILL_200)
		case 10:
			depositMoney(w, BYN_BILL_500)
		default:
			showWrongInputMessage()
		}
	}
}
