package services

import (
	"bytes"
	"cart/dotEnv"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type DefaultResponse struct {
	Status  string
	Message string
}

func RemoveProductFromCart(productBody Product) error {
	body, _ := json.Marshal(productBody)
	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("REMOVE_PRODUCT_FROM_CART")

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}

	res, clientErr := client.Do(req)
	if clientErr != nil {
		return clientErr
	}

	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}
	defer res.Body.Close()

	var resp DefaultResponse
	if unmarshalErr := json.Unmarshal([]byte(b), &resp); unmarshalErr != nil {
		return unmarshalErr
	}

	if resp.Message == "" {
		return errors.New("response empty")
	}

	return nil
}
