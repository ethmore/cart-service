package services

import (
	"bytes"
	"cart-service/dotEnv"
	"encoding/json"
	"errors"
	"io"
	"net/http"
	"time"
)

type PurchaseBody struct {
	Token      string
	TotalPrice string
}

func AddToCart(product Product) error {
	body, _ := json.Marshal(product)
	bodyReader := bytes.NewReader(body)
	requestURL := dotEnv.GoDotEnvVariable("ADD_PRODUCT_TO_CART")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
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

	//response
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}
	defer res.Body.Close()

	var resp ResponseBody
	if unmarshalErr := json.Unmarshal([]byte(b), &resp); unmarshalErr != nil {
		return unmarshalErr
	}

	if resp.Message == "" {
		return errors.New("response empty")
	}

	return nil
}

func AddTotalToCart(purchaseBody PurchaseBody) error {
	body, _ := json.Marshal(purchaseBody)
	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("ADD_TOTAL_TO_CART")

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

	_, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}
	defer res.Body.Close()

	return nil
}
