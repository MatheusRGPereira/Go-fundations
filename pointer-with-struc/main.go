package main

type Account struct {
	balance int
}

func (c *Account) Simule(value int) int {
	c.balance += value
	println(c.balance)
	return c.balance
}

func newAccount() *Account {
	return &Account{
		balance: 0,
	}
}

func main() {
	Account := Account{
		balance: 1500,
	}

	Account.Simule(500.00)

	println(Account.balance)

}
