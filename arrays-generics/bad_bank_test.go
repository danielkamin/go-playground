package arraysgenerics

import "testing"

func TestBadBank(t *testing.T) {
	var (
		riya  = Account{Name: "Ryiya", Balance: 100}
		chris = Account{Name: "Ryiya", Balance: 75}
		adil  = Account{Name: "Adil", Balance: 200}
	)
	transactions := []Transaction{
		{
			From: "Chris",
			To:   "Riya",
			Sum:  100,
		},
		{
			From: "Adil",
			To:   "Chris",
			Sum:  25,
		},
	}
	newBalanceFor := func(account Account) float64 {
		return NewBalanceFor(account, transactions).Balance
	}
	AssertEqual(t, newBalanceFor(riya), 100.0)
	AssertEqual(t, newBalanceFor(chris), 75.0)
	AssertEqual(t, newBalanceFor(adil), 175.0)
}
