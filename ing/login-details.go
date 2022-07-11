package ing

import "github.com/alshdavid/open-bank-api/bank"

func NewLoginDetails(clientId string, pin string) bank.ILoginDetails {
	return bank.NewLoginDetails(map[string]string{
		Meta.ClientID: clientId,
		Meta.Pin:      pin,
	})
}
