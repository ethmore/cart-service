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

func ChangeProductQty(product Product) error {
	body, _ := json.Marshal(product)
	bodyReader := bytes.NewReader(body)
	requestURL := dotEnv.GoDotEnvVariable("CHANGE_PRODUCT_QTY")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		return err
	}

	//response
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return readErr
	}
	defer res.Body.Close()

	var resp ResponseBody
	if err := json.Unmarshal([]byte(b), &resp); err != nil {
		return err
	}

	if resp.Message == "" {
		return errors.New("response empty")
	}

	return nil
}
