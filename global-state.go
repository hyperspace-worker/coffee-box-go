package main

type GlobalState struct {
	wallet           Wallet
	storage          ItemStorage
	pinInputAttempts int
}
