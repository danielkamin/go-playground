package arraysgenerics

type Transaction struct {
	From, To string
	Sum      float64
}
type Account struct {
	Name    string
	Balance float64
}

func BalanceFor(transactions []Transaction, who string) float64 {
	adjustbalance := func(currentBalance float64, t Transaction) float64 {
		if t.From == who {
			return currentBalance - t.Sum
		}
		if t.To == who {
			return currentBalance + t.Sum
		}
		return currentBalance
	}

	return Reduce(transactions, adjustbalance, 0.0)
}
func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}
func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransactions,
		account,
	)
}

func applyTransactions(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}
	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}
	return a
}
