package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

const ROW_LENGTH int = 40
const PRICE_CAPPUCCINO float32 = 3
const PRICE_ESPRESSO float32 = 3.5
const PRICE_LATTE float32 = 4
const GLASSES_CAPACITY int = 600
const PIN int = 7878

const BYN_BILL_05 float32 = 0.5
const BYN_BILL_1 float32 = 1
const BYN_BILL_2 float32 = 2
const BYN_BILL_5 float32 = 5
const BYN_BILL_10 float32 = 10
const BYN_BILL_20 float32 = 20
const BYN_BILL_50 float32 = 50
const BYN_BILL_100 float32 = 100
const BYN_BILL_200 float32 = 200
const BYN_BILL_500 float32 = 500

const MANAGER_CONTACTS string = "+375 (29) 197-15-64"
const MAX_PIN_INPUT_ATTEMPTS int = 3

func main() {

	fmt.Println("Coffe machine is starting...")
	fmt.Println("Firmware version is 4.0.2")
	fmt.Println("Firmware is up to date")
	fmt.Println("Loading...")

	time.Sleep(2 * time.Second)

	fmt.Println("Coffe machine is ready!")

	time.Sleep(2 * time.Second)

	var glasses int = 7
	var availablePinInputAttempts int = 3

	var userBalance float32 = 0
	var cashBalance float32 = 0

	clearScreen()

	callMainMenu(glasses, userBalance, cashBalance, availablePinInputAttempts)
}

/******************************************************************************
 *
 * MAIN MENU FUNCTIONS
 *
******************************************************************************/

func callMainMenu(glasses int, userBalance float32, cashBalance float32, availablePinInputAttempts int) {

	r := bufio.NewReader(os.Stdin)

	for true {
		choiseOption := 0

		showMainMenu(userBalance, glasses)

		fmt.Print("Please, choise option: ")

		res, err := selectOption(r)

		fmt.Printf("Selected option is %v, err is %v\n", res, err)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			time.Sleep(2 * time.Second)
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
				showSymbolsRow()
				showSymbolsRowWithMessage("You cancelled operation", ROW_LENGTH)
				showSymbolsRow()
				fmt.Println()
				fmt.Println()
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
	showLogo()
	showSymbolsRowWithMessage("MAIN MENU", ROW_LENGTH)
	showSymbolsRow()
	fmt.Printf("%-25v %v BYN\n", "Cash balance:", userBalance)
	showSymbolsRow()
	fmt.Printf("%-25v %v\n", "Number of glasses:", glasses)
	showSymbolsRow()
	showCoffeeList()
	fmt.Printf("*%-24v %-15v", "4. Cash deposit", "*")
	showSymbolsRow()
	fmt.Printf("*%-19v %-20v", "5. Service", "*")
	showSymbolsRow()

}

func showLogo() {
	showHeader("ESPRESSO BIANCCI")
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

func showCoffeeList() {
	showSymbolsRowWithMessage("Select coffee", ROW_LENGTH)
	showSymbolsRow()
	fmt.Printf("*%-22v %-5v %-12v\n", "1. Cappuccino", PRICE_CAPPUCCINO, "*")
	fmt.Printf("*%-20v %-7v %-12v\n", "2. Espresso", PRICE_ESPRESSO, "*")
	fmt.Printf("*%-17v %-10v %-12v\n", "3. Latte", PRICE_LATTE, "*")
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
	showSymbolsRow()
	showSymbolsRowWithMessage("Take your coffee, please!", ROW_LENGTH)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}

func showNotEnoughMoneyWarning() {
	showSymbolsRow()
	showSymbolsRowWithMessage("Sorry, you don't have enough money!", ROW_LENGTH)
	showSymbolsRowWithMessage("You need to put some cash", ROW_LENGTH)
	showSymbolsRowWithMessage("in coffe machine", ROW_LENGTH)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}

func showNoGlassesWarning() {
	showSymbolsRow()
	showSymbolsRowWithMessage("Sorry we don't have glasses! Please,", ROW_LENGTH)
	showSymbolsRowWithMessage("call our manager: ", ROW_LENGTH)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}

func showMoneyFromUser(byn float32) {
	showSymbolsRow()
	fmt.Printf(" You put in coffe machine %0.1f BYN \n", byn)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}

func showSymbolsRow() {
	for i := 0; i < 40; i++ {
		fmt.Print("*")
	}
	fmt.Print("\n")
}

func showSymbolsRowWithMessage(message string, rowLength int) {
	fmt.Print("*")

	diff := rowLength - len(message)

	for i := 0; i < (diff / 2); i++ {
		fmt.Print(" ")
	}

	fmt.Print(message)

	if diff%2 == 1 {
		for i := 0; i < diff; i++ {
			fmt.Print(" ")
		}
	} else {
		for i := 0; i < (diff/2)-1; i++ {
			fmt.Print(" ")
		}
	}

	fmt.Print("*\n")
}

func showWrongInputMessage() {
	showSymbolsRow()
	showSymbolsRowWithMessage("Wrong input! Try again...", ROW_LENGTH)
	showSymbolsRow()

	fmt.Println()
	fmt.Println()
}

func callCashDepositMenu(userBalance float32, cashBalance float32) int {

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
			time.Sleep(2 * time.Second)
			clearScreen()
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

/******************************************************************************
 *
 * SERVICE MENU FUNCTIONS
 *
******************************************************************************/

func callServiceMenu(currentGlassesNumber int, allowedCash float32, currentUserBalance float32, availablePinInputAttempts int) int {

	r := bufio.NewReader(os.Stdin)

	for {
		choiseOption := -1
		availablePinInputAttempts = MAX_PIN_INPUT_ATTEMPTS

		showServiceMenu(currentGlassesNumber, allowedCash)

		fmt.Println("Select option or press 0 to exit: ")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			time.Sleep(2 * time.Second)
			clearScreen()
			continue
		}

		choiseOption = res

		clearScreen()

		switch choiseOption {
		case 0:
			currentUserBalance = 0
			return 0
		case 1:
			giveOutProceed(allowedCash)
		case 2:
			fillCoffeeMachineWithGlasses(currentGlassesNumber)
		default:
			showWrongInputMessage()
		}
	}
}

func showServiceMenu(currentGlassesNubmer int, allowedCash float32) {
	showHeader("Service menu")
	fmt.Printf("%-25v %v BYN\n", "Cash balance:", allowedCash)
	showSymbolsRow()
	fmt.Printf("%-25v %v BYN\n", "Glasses left:", currentGlassesNubmer)
	showSymbolsRow()
	showSymbolsRowWithMessage("1. Issue proceeds", ROW_LENGTH)
	showSymbolsRowWithMessage("2. Load the glasses", ROW_LENGTH)
	showSymbolsRowWithMessage("0. Exit", ROW_LENGTH)
	showSymbolsRow()
}

func showProceeds(cash float32) {
	fmt.Printf("Available proceeds is %v BYN", cash)
}

func showGlasses(glassesCount int) {
	fmt.Printf("%v glasses left", glassesCount)
}

func giveOutProceed(availableCash float32) {
	availableCash = 0

	showSymbolsRow()
	showSymbolsRowWithMessage("Opening the lock...", ROW_LENGTH)
	showSymbolsRowWithMessage("Opened.", ROW_LENGTH)
	showSymbolsRowWithMessage("You successfully take out", ROW_LENGTH)
	showSymbolsRowWithMessage("all proceeds.", ROW_LENGTH)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}

/******************************************************************************
 *
 * GENERAL FUNCTIONS
 *
******************************************************************************/

// TODO Not a cross-platform solution...
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func fillCoffeeMachineWithGlasses(glassesLeft int) {
	r := bufio.NewReader(os.Stdin)
	newGlasses := 0
	totalGlasses := glassesLeft
	leftCapacity := GLASSES_CAPACITY - glassesLeft

	fmt.Println("How many glasses you want to insert?")
	res, err := selectOption(r)

	if err != nil {
		clearScreen()
		showWrongInputMessage()
		time.Sleep(2 * time.Second)
		clearScreen()
		return
	}

	newGlasses = res

	clearScreen()
	showSymbolsRow()

	totalGlasses += newGlasses

	if newGlasses <= 0 {
		showSymbolsRowWithMessage("Are you kidding me?", ROW_LENGTH)
	} else if totalGlasses > GLASSES_CAPACITY {
		showSymbolsRowWithMessage("Too much glasses!", ROW_LENGTH)

		if leftCapacity > 0 {
			showSymbolsRowWithMessage("Try to insert less. You can only", ROW_LENGTH)
			showSymbolsRowWithMessage(fmt.Sprintf("load %v glasses.", leftCapacity), ROW_LENGTH)
		} else {
			showSymbolsRowWithMessage("You can't load any glasses!", ROW_LENGTH)
			showSymbolsRowWithMessage("Container is full!", ROW_LENGTH)
		}
	} else {
		glassesLeft = totalGlasses
		showSymbolsRowWithMessage("You successfully filled up the coffee box", ROW_LENGTH)
		showSymbolsRowWithMessage(fmt.Sprintf("with %v glasses", newGlasses), ROW_LENGTH)
	}
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}

func checkAccess(availablePinInputAttempts int) int {
	userChoise := 0
	r := bufio.NewReader(os.Stdin)

	for availablePinInputAttempts > 0 {
		fmt.Println("Please, enter a PIN number")
		fmt.Print("or press 0 to exit: ")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			time.Sleep(2 * time.Second)
			clearScreen()
		}

		userChoise = res

		clearScreen()

		if userChoise == 0 {
			return 0
		}
		if userChoise == PIN {
			return 1
		}

		availablePinInputAttempts--
		showSymbolsRow()
		showSymbolsRowWithMessage("Wrong PIN number!", ROW_LENGTH)
		showSymbolsRow()
		fmt.Println()
	}

	return -1
}

func checkGlasses(availableGlasses int) int {
	if availableGlasses == 0 {
		return 0
	}
	return availableGlasses
}

func isMoneyEnough(currentBalance float32, itemPrice float32) bool {
	return currentBalance >= itemPrice
}

func addSugar() int {
	choiseOption := 0
	r := bufio.NewReader(os.Stdin)

	for {
		showSymbolsRow()
		showSymbolsRowWithMessage("Would you like to add sugar?", ROW_LENGTH)
		showSymbolsRow()
		showSymbolsRowWithMessage("  Yes - 1        No - 0     ", ROW_LENGTH)
		showSymbolsRow()
		fmt.Println("Please, choise option:")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			time.Sleep(2 * time.Second)
			clearScreen()
			continue
		}

		choiseOption = res

		switch choiseOption {
		case 0:
			return 0
		case 1:
			adjustPortionSize()
			return 0
		default:
			showWrongInputMessage()
			fmt.Println()
		}
	}
}

func adjustPortionSize() int {
	currentPortionSize := 4
	choiseOption := -1
	r := bufio.NewReader(os.Stdin)

	for {
		showSymbolsRow()
		showSymbolsRowWithMessage("0 - confirm", ROW_LENGTH)
		showSymbolsRowWithMessage("1 - increase", ROW_LENGTH)
		showSymbolsRowWithMessage("2 - decrease", ROW_LENGTH)
		showSymbolsRow()
		fmt.Println("Please, choise option:")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
			time.Sleep(2 * time.Second)
			clearScreen()
			continue
		}

		choiseOption = res

		switch choiseOption {
		case 0:
			return currentPortionSize
		case 1:
			if currentPortionSize >= 8 {
				showSymbolsRow()
				showSymbolsRowWithMessage("Sorry, can't increase sugar portion!", ROW_LENGTH)
				showSymbolsRowWithMessage("It's maximum", ROW_LENGTH)
				showSymbolsRow()
				fmt.Println()
				fmt.Println()
			} else {
				currentPortionSize++
			}
		case 2:
			if currentPortionSize <= 0 {
				showSymbolsRow()
				showSymbolsRowWithMessage("Sorry, can't decrease sugar portion!", ROW_LENGTH)
				showSymbolsRowWithMessage("It's minimum!", ROW_LENGTH)
				showSymbolsRow()
				fmt.Println()
				fmt.Println()
			} else {
				currentPortionSize--
			}
		default:
			showWrongInputMessage()
		}
	}
}

func giveCoffeeToUser(userBalance float32, price float32, glasses int) {
	if checkGlasses(glasses) > 0 {
		if isMoneyEnough(userBalance, price) {
			addSugar()
			showCoffeeIsPurchased()
			userBalance -= price
			glasses--
		} else {
			showNotEnoughMoneyWarning()
		}
	} else {
		showNoGlassesWarning()
	}
}

func getMoneyFromUser(userBalance float32, cashBalance float32, byn float32) {
	showMoneyFromUser(byn)
	userBalance += byn
	cashBalance += byn
}

func selectOption(r *bufio.Reader) (int, error) {
	val, err := r.ReadString('\n')

	if err != nil {
		fmt.Println("error reading string")
		return -1, err
	}

	val = strings.TrimRight(val, "\r\n")
	res, err := strconv.Atoi(val)

	if err != nil {
		fmt.Println("error parsing value")
		return -1, err
	}

	if res < 0 {
		return -1, fmt.Errorf("Option must be a positive number! (current value %v)", res)
	}

	return res, nil
}
