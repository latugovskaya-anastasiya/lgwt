package pointersErrors

import "fmt"

type Stringer interface {
	String() string
}

type Wallet struct {
	balance Bitcoin
}

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

func (w *Wallet) Deposit(amount Bitcoin) {
	// fmt.Printf("address of balance in Deposit is %v \n", &w.balance)
	w.balance += amount
}
func (w *Wallet) Balance() Bitcoin {
	return w.balance // or (*w.)balance
}