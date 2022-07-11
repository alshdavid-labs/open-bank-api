package bank

type ISession interface {
	// GetAccounts will return a list of accounts held by this bank
	GetAccounts() string
}
