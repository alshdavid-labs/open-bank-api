package ingaus

import (
	"errors"
	"fmt"

	"github.com/alshdavid/open-bank-api/bank"
	"github.com/alshdavid/open-bank-api/ing-aus/internal/pinpad"
	"github.com/alshdavid/open-bank-api/ing-aus/internal/requests"
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

	pinPadResponse := requests.GetPinPadImages()

	for pos, b64Image := range pinPadResponse.KeypadImages {
		digitImage, err := pinpad.Base64ToRGBA(b64Image)
		if err != nil {
			fmt.Println("Error decoding b64")
			continue
		}

		digit, _ := pinpad.FindDigit(digitImage)

		fmt.Printf("Digit at pos %d is %d\n", pos, digit)
	}

	return &Session{}, nil
}
