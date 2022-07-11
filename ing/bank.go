package ing

import (
	"errors"

	"github.com/alshdavid/open-bank-api/bank"
)

type Bank struct{}

func NewBank() bank.IBank {
	return &Bank{}
}

func (*Bank) GetLoginMeta() bank.ILoginMeta {
	return LoginMeta
}

func (*Bank) Login(details bank.ILoginDetails) (bank.ISession, error) {
	clientId := details.Get(Meta.ClientID)
	pin := details.Get(Meta.Pin)

	if clientId == "" || pin == "" {
		return nil, errors.New(ErrorType.InvalidCredentials)
	}

	return &Session{}, nil
}
