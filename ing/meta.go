package ing

import "github.com/alshdavid/open-bank-api/bank"

const (
	MetaClientID = "ClientID"
	MetaPin      = "Pin"
)

var Meta = struct {
	ClientID string
	Pin      string
}{
	ClientID: MetaClientID,
	Pin:      MetaPin,
}

var MetaDescription = struct {
	ClientID string
	Pin      string
}{
	ClientID: "Account identifier",
	Pin:      "Numeric pass code used to authenticate",
}

var LoginMeta = bank.NewLoginMeta(map[string]string{
	Meta.ClientID: MetaDescription.ClientID,
	Meta.Pin:      MetaDescription.Pin,
})
