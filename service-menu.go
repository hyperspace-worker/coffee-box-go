package main

import (
	"bufio"
	"fmt"
	"os"
)

func callServiceMenu(gs *GlobalState) int {

	r := bufio.NewReader(os.Stdin)

	for {
		choiseOption := -1

		showServiceMenu(gs)

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
			gs.wallet.resetUserBalance()
			return 0
		case 1:
			giveOutProceed(&gs.wallet)
		case 2:
			fillCoffeeMachineWithGlasses(&gs.storage)
		default:
			showWrongInputMessage()
		}
	}
}

func showServiceMenu(gs *GlobalState) {
	showHeader("Service menu")
	fmt.Printf("%-25v %v BYN\n", "Cash balance:", gs.wallet.proceeds)
	showSymbolsRow()
	fmt.Printf("%-25v %v BYN\n", "Glasses left:", gs.storage.cups)
	showSymbolsRow()
	showSymbolsRowWithMessage("1. Issue proceeds", ROW_LENGTH)
	showSymbolsRowWithMessage("2. Load the glasses", ROW_LENGTH)
	showSymbolsRowWithMessage("0. Exit", ROW_LENGTH)
	showSymbolsRow()
}

func giveOutProceed(w *Wallet) {
	proceeds := w.withdrawProceeds()

	showSymbolsRow()
	showSymbolsRowWithMessage("Opening the lock...", ROW_LENGTH)
	showSymbolsRowWithMessage("Opened.", ROW_LENGTH)
	showSymbolsRowWithMessage("You successfully take out", ROW_LENGTH)
	showSymbolsRowWithMessage("all proceeds", ROW_LENGTH)
	showSymbolsRowWithMessage(fmt.Sprintf("(%0.2f BYN total)", proceeds), ROW_LENGTH)
	showSymbolsRow()
	fmt.Println()
	fmt.Println()
}
