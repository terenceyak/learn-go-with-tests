package wallet

import "fmt"

type Stringer interface {
	String() string
}

type Bitcoin int

type Wallet struct {
	balance Bitcoin
}

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	fmt.Printf("address of balance in test is %v \n", &w.balance)
	w.balance += amount
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}
