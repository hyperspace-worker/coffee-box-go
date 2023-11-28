package main

type Wallet struct {
	balance  float32
	proceeds float32
}

func (w *Wallet) withdrawProceeds() float32 {
	proceeds := w.proceeds
	w.proceeds = 0
	return proceeds
}

func (w *Wallet) resetUserBalance() float32 {
	currentBalance := w.balance
	w.balance = 0
	return currentBalance
}

func (w *Wallet) depositMoney(m float32) float32 {
	w.balance += m
	w.proceeds += m
	return m
}

func (w *Wallet) tryWithdrawMoney(itemPrice float32) bool {
	if itemPrice > w.balance {
		return false
	}
	w.balance -= itemPrice
	return true
}
