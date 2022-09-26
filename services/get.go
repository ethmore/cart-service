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

type Product struct {
	Token string
	Id    string
	Qty   string
}

type ProductResponse struct {
	Id  string
	Qty string
}

type ResponseBody struct {
	Message  string
	Products []ProductResponse
}

type ProductInfo struct {
	Id          string
	Title       string
	Price       string
	Description string
	Image       string
	Stock       string
	Qty         string
}

type GetProductResponse struct {
	Products ProductInfo `json:"product"`
}

func GetCartProducts(responseBody Product) (string, []ProductInfo, error) {
	body, _ := json.Marshal(responseBody)
	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("GET_CART_PRODUCTS")

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		return "", nil, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, clientErr := client.Do(req)
	if clientErr != nil {
		return "", nil, clientErr
	}

	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return "", nil, readErr
	}
	defer res.Body.Close()

	var resp ResponseBody
	if respErr := json.Unmarshal([]byte(b), &resp); respErr != nil {
		return "", nil, respErr
	}

	if resp.Message == "" {
		return "", nil, errors.New("response empty")
	}

	products := []ProductInfo{}
	for i := 0; i < len(resp.Products); i++ {
		product, pGetErr := getProductInfo(responseBody.Token, resp.Products[i].Id)
		product.Qty = resp.Products[i].Qty
		if pGetErr != nil {
			return "", nil, pGetErr
		}
		products = append(products, *product)
	}

	return resp.Message, products, nil
}

func getProductInfo(token, productId string) (*ProductInfo, error) {
	bo := Product{
		Token: token,
		Id:    productId,
	}
	body, _ := json.Marshal(bo)
	bodyReader := bytes.NewReader(body)
	requestUrl := dotEnv.GoDotEnvVariable("GET_PRODUCT")

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, clientErr := client.Do(req)
	if clientErr != nil {
		return nil, clientErr
	}

	//response
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return nil, err
	}
	defer res.Body.Close()

	var resp GetProductResponse
	if unmarshalErr := json.Unmarshal([]byte(b), &resp); unmarshalErr != nil {
		return nil, unmarshalErr
	}

	if resp.Products.Id == "" {
		return nil, errors.New("response empty")
	}
	return &resp.Products, nil
}

func GetTotalPrice(response PurchaseBody) (totalPrice string, e error) {
	body, _ := json.Marshal(response)
	bodyReader := bytes.NewReader(body)
	requestURL := dotEnv.GoDotEnvVariable("GET_TOTAL_PRICE")

	req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
	if err != nil {
		return "", err
	}
	req.Header.Set("Content-Type", "application/json")

	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, clientErr := client.Do(req)
	if clientErr != nil {
		return "", clientErr
	}

	//response
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		return "", readErr
	}
	defer res.Body.Close()

	var resp PurchaseBody
	if unmarshalErr := json.Unmarshal([]byte(b), &resp); unmarshalErr != nil {
		return "", unmarshalErr
	}

	return resp.TotalPrice, nil
}
