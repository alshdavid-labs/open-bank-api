package requests

type GetPinPadImagesResponse struct {
	PemEncryptionKey string   `json"PemEncryptionKey"`
	Secret           string   `json"Secret"`
	KeypadImages     []string `json"KeypadImages"`
}

func GetPinPadImages() {}
