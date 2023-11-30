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

	wallet := Wallet{}
	storage := ItemStorage{cups: 7}
	state := GlobalState{wallet, storage, MAX_PIN_INPUT_ATTEMPTS}

	clearScreen()

	callMainMenu(&state)
}
