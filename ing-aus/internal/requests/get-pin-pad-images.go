package requests

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type GetPinPadImagesResponse struct {
	PemEncryptionKey string   `json"PemEncryptionKey"`
	Secret           string   `json"Secret"`
	KeypadImages     []string `json"KeypadImages"`
}

func GetPinPadImages() *GetPinPadImagesResponse {
	resp, err := http.Get(URLs.GetPinPadImages())
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	response := &GetPinPadImagesResponse{}
	json.Unmarshal(body, response)

	return response
}
