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

// TODO Not a cross-platform solution...
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func fillCoffeeMachineWithGlasses(glassesLeft *int) {
	r := bufio.NewReader(os.Stdin)
	newGlasses := 0
	totalGlasses := *glassesLeft
	leftCapacity := GLASSES_CAPACITY - *glassesLeft

	fmt.Println("How many glasses you want to insert?")
	res, err := selectOption(r)

	if err != nil {
		clearScreen()
		showWrongInputMessage()
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
		*glassesLeft = totalGlasses
		showSymbolsRowWithMessage("You successfully filled up the coffee box", ROW_LENGTH)
		showSymbolsRowWithMessage(fmt.Sprintf("with %v glasses", newGlasses), ROW_LENGTH)
	}
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
	time.Sleep(2 * time.Second)
	clearScreen()
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

func isMoneyEnough(currentBalance *float32, itemPrice float32) bool {
	return *currentBalance >= itemPrice
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
	currentPortionSize := MAX_SUGAR_PORTION_SIZE / 2
	choiseOption := -1
	r := bufio.NewReader(os.Stdin)

	for {
		clearScreen()
		showSymbolsRow()

		for i := 0; i < currentPortionSize; i++ {
			fmt.Print("▓", " ")
		}

		sdf := MAX_SUGAR_PORTION_SIZE - currentPortionSize

		for i := 0; i < sdf; i++ {
			fmt.Print("░", " ")
		}

		fmt.Print("\n")

		showSymbolsRowWithMessage("0 - confirm", ROW_LENGTH)
		showSymbolsRowWithMessage("1 - increase", ROW_LENGTH)
		showSymbolsRowWithMessage("2 - decrease", ROW_LENGTH)
		showSymbolsRow()
		fmt.Println("Please, choise option:")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
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
				time.Sleep(2 * time.Second)
			} else {
				currentPortionSize++
			}
		case 2:
			if currentPortionSize <= 0 {
				showSymbolsRow()
				showSymbolsRowWithMessage("Sorry, can't decrease sugar portion!", ROW_LENGTH)
				showSymbolsRowWithMessage("It's minimum!", ROW_LENGTH)
				showSymbolsRow()
				time.Sleep(2 * time.Second)
			} else {
				currentPortionSize--
			}
		default:
			showWrongInputMessage()
		}

		clearScreen()
	}
}

func giveCoffeeToUser(userBalance *float32, price float32, glasses *int) {
	if checkGlasses(*glasses) > 0 {
		if isMoneyEnough(userBalance, price) {
			addSugar()
			*userBalance -= price
			*glasses--
			showCoffeeIsPurchased()
		} else {
			showNotEnoughMoneyWarning()
		}
	} else {
		showNoGlassesWarning()
	}
}

func getMoneyFromUser(userBalance *float32, cashBalance *float32, byn float32) {
	showMoneyFromUser(byn)
	*userBalance += byn
	*cashBalance += byn
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
