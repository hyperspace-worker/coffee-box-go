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

func showMessage(message string) {
	showSymbolsRow()
	showSymbolsRowWithMessage(message, ROW_LENGTH)
	showSymbolsRow()
	time.Sleep(time.Duration(MESSAGE_DISPLAY_DURATION) * time.Second)
	clearScreen()
}

func showWrongInputMessage() {
	showMessage("Wrong input! Try again...")
}

// TODO Not a cross-platform solution...
func clearScreen() {
	cmd := exec.Command("cmd", "/c", "cls")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func fillCoffeeMachineWithGlasses(s *ItemStorage) {
	r := bufio.NewReader(os.Stdin)
	newGlasses := 0
	totalGlasses := s.cups
	leftCapacity := GLASSES_CAPACITY - s.cups

	fmt.Println("How many glasses you want to insert?")
	res, err := selectOption(r)

	if err != nil {
		clearScreen()
		showWrongInputMessage()
		return
	}

	newGlasses = res

	clearScreen()
	showSymbolsRow()

	totalGlasses += newGlasses

	if newGlasses <= 0 {
		showMessage("Are you kidding me?")
	} else if totalGlasses > GLASSES_CAPACITY {
		if leftCapacity > 0 {
			showMessage(fmt.Sprintf("Too much glasses! Try to insert less. You can only load %v glasses", leftCapacity))
		} else {
			showMessage("Too much glasses! You can't load any glasses! Container is full!")
		}
	} else {
		s.cups = totalGlasses
		showMessage(fmt.Sprintf("You successfully filled up the coffee box with %v glasses", newGlasses))
	}
}

func checkAccess(availablePinInputAttempts int) int {
	userChoise := 0
	r := bufio.NewReader(os.Stdin)

	for availablePinInputAttempts > 0 {
		fmt.Println("Please, enter a PIN number or press 0 to exit:")

		res, err := selectOption(r)

		if err != nil {
			clearScreen()
			showWrongInputMessage()
		}

		userChoise = res

		clearScreen()

		if userChoise == 0 {
			return 0
		}
		if userChoise == PIN {
			return 1
		}

		// TODO Do we need to restore pin input attempts if user finally successfully entered service menu?
		availablePinInputAttempts--

		showMessage("Wrong PIN number")
	}

	return -1
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
			continue
		}

		choiseOption = res

		switch choiseOption {
		case 0:
			return currentPortionSize
		case 1:
			if currentPortionSize >= 8 {
				showMessage("Sorry, can't increase sugar portion! It's maximum")
			} else {
				currentPortionSize++
			}
		case 2:
			if currentPortionSize <= 0 {
				showMessage("Sorry, can't decrease sugar portion! It's minimum!")
			} else {
				currentPortionSize--
			}
		default:
			showWrongInputMessage()
		}

		clearScreen()
	}
}

// TODO Add notification if glasses count less than 5 pc
func giveCoffeeToUser(w *Wallet, s *ItemStorage, price float32) {
	if !s.areAnyGlasses() {
		showNoGlassesWarning()
		return
	}

	isSuccess := w.tryWithdrawMoney(price)

	if isSuccess {
		addSugar()
		s.getCup()
		showCoffeeIsPurchased()
		return
	}

	showNotEnoughMoneyWarning()
}

func depositMoney(w *Wallet, byn float32) {
	w.depositMoney(byn)
	showMessage(fmt.Sprintf(" You put in coffe machine %0.2f BYN \n", byn))
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
