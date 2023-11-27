package main

import (
	"fmt"
	"time"
)

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

	callMainMenu(&glasses, &userBalance, &cashBalance, availablePinInputAttempts)
}
