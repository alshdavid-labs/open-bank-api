package requests

var URLs = struct {
	GetPinPadImages func() string
}{
	GetPinPadImages: func() string {
		return "https://www.ing.com.au/KeypadService/v1/KeypadService.svc/json/PinpadImages"
	},
}
