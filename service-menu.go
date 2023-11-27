package main

import (
	"bufio"
	"fmt"
	"os"
)

func callServiceMenu(currentGlassesNumber *int, allowedCash *float32, currentUserBalance *float32, availablePinInputAttempts int) int {

	r := bufio.NewReader(os.Stdin)

	for {
		choiseOption := -1
		availablePinInputAttempts = MAX_PIN_INPUT_ATTEMPTS

		showServiceMenu(*currentGlassesNumber, *allowedCash)

		fmt.Println("Select option or press 0 to exit: ")

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
			*currentUserBalance = 0
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

func giveOutProceed(availableCash *float32) {
	*availableCash = 0

	showSymbolsRow()
	showSymbolsRowWithMessage("Opening the lock...", ROW_LENGTH)
	showSymbolsRowWithMessage("Opened.", ROW_LENGTH)
	showSymbolsRowWithMessage("You successfully take out", ROW_LENGTH)
	showSymbolsRowWithMessage("all proceeds.", ROW_LENGTH)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}
