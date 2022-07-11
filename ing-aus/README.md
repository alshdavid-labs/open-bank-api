# ING Australia API Wrapper

## Usage

```go
package main

import (
	"fmt"

	"github.com/alshdavid/open-bank-api/bank"
	ing_aus "github.com/alshdavid/open-bank-api/ing-aus"
)

func main() {
	var ingBank bank.IBank = ing_aus.NewBank()

	// Log into bank with your credentials
	loginDetails := ing_aus.NewLoginDetails("ClientID", "ClientPin")
	session, _ := ingBank.Login(loginDetails)

	// Print out your accounts
	fmt.Println(session.GetAccounts())
}
```