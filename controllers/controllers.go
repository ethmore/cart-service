package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

type DefaultResponse struct {
	Status  string
	Message string
}

type Product struct {
	Token string
	Id    string
	Qty   string
}

type ResponseBody struct {
	Message  string
	Products []ProductResponse
}

type ProductResponse struct {
	Id  string
	Qty string
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
	Message  string
	Products ProductInfo
}

type PurchaseBody struct {
	Token      string
	TotalPrice string
}

func AddToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Product
		if bodyErr := ctx.ShouldBindBodyWith(&product, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		body, _ := json.Marshal(product)
		bodyReader := bytes.NewReader(body)
		requestURL := "http://127.0.0.1:3002/addProductToCart"

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		//response
		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		var resp ResponseBody
		json.Unmarshal([]byte(b), &resp)

		if resp.Message == "" {
			fmt.Println("response empty")
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetCartProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var responseBody Product
		if bodyErr := ctx.ShouldBindBodyWith(&responseBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		body, _ := json.Marshal(responseBody)
		bodyReader := bytes.NewReader(body)
		requestUrl := "http://127.0.0.1:3002/getCartProducts"

		req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
		if err != nil {
			fmt.Println("client could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client error making http request: ", err)
			return
		}

		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		var resp ResponseBody
		json.Unmarshal([]byte(b), &resp)
		if resp.Message == "" {
			fmt.Println("response empty")
			ctx.Status(http.StatusInternalServerError)
			return
		}

		products := []ProductInfo{}
		for i := 0; i < len(resp.Products); i++ {
			product, err := GetProductInfo(responseBody.Token, resp.Products[i].Id)
			product.Qty = resp.Products[i].Qty
			if err != nil {
				fmt.Println("product info:", err)
				ctx.Status(http.StatusInternalServerError)
				return
			}
			products = append(products, *product)
		}
		fmt.Println(products)
		ctx.JSON(200, gin.H{"message": resp.Message, "products": products})
	}
}

func GetProductInfo(token, productId string) (*ProductInfo, error) {
	bo := Product{
		Token: token,
		Id:    productId,
	}
	body, _ := json.Marshal(bo)
	bodyReader := bytes.NewReader(body)
	requestUrl := "http://127.0.0.1:3002/getProduct"

	req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
	if err != nil {
		fmt.Println("client: could not create request", err)
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	client := http.Client{
		Timeout: 30 * time.Second,
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println("client: error making http request: ", err)
		return nil, err
	}

	//response
	b, readErr := io.ReadAll(res.Body)
	if readErr != nil {
		fmt.Println(readErr)
		return nil, err
	}
	defer res.Body.Close()
	var resp GetProductResponse
	json.Unmarshal([]byte(b), &resp)

	if resp.Products.Id == "" {
		fmt.Println("response empty")
		return nil, errors.New("response empty")
	}
	return &resp.Products, nil
}

func RemoveProductFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var fBody Product
		if bodyErr := ctx.ShouldBindBodyWith(&fBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		body, _ := json.Marshal(fBody)
		bodyReader := bytes.NewReader(body)
		requestUrl := "http://127.0.0.1:3002/removeProductFromCart"

		req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}

		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		var resp DefaultResponse
		json.Unmarshal([]byte(b), &resp)

		if resp.Message == "" {
			fmt.Println("response empty")
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func IncreaseProductQty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Product
		if bodyErr := ctx.ShouldBindBodyWith(&product, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		body, _ := json.Marshal(product)
		bodyReader := bytes.NewReader(body)
		requestURL := "http://127.0.0.1:3002/increaseProductQty"

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		//response
		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		var resp ResponseBody
		json.Unmarshal([]byte(b), &resp)

		if resp.Message == "" {
			fmt.Println("response empty")
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func DecreaseProductQty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product Product
		if bodyErr := ctx.ShouldBindBodyWith(&product, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		body, _ := json.Marshal(product)
		bodyReader := bytes.NewReader(body)
		requestURL := "http://127.0.0.1:3002/decreaseProductQty"

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		//response
		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		var resp ResponseBody
		json.Unmarshal([]byte(b), &resp)

		if resp.Message == "" {
			fmt.Println("response empty")
			ctx.Status(http.StatusInternalServerError)
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func ToPurchase() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var purchaseBody PurchaseBody
		if bodyErr := ctx.ShouldBindBodyWith(&purchaseBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			return
		}

		body, _ := json.Marshal(purchaseBody)
		bodyReader := bytes.NewReader(body)
		requestUrl := "http://127.0.0.1:3002/addTotalToCart"

		req, err := http.NewRequest(http.MethodPost, requestUrl, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		fmt.Println(b)
		//...

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetTotalPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response PurchaseBody
		if bodyErr := ctx.ShouldBindBodyWith(&response, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		body, _ := json.Marshal(response)
		bodyReader := bytes.NewReader(body)
		requestURL := "http://127.0.0.1:3002/getTotalPrice"

		req, err := http.NewRequest(http.MethodPost, requestURL, bodyReader)
		if err != nil {
			fmt.Println("client: could not create request", err)
			return
		}
		req.Header.Set("Content-Type", "application/json")

		client := http.Client{
			Timeout: 30 * time.Second,
		}
		res, err := client.Do(req)
		if err != nil {
			fmt.Println("client: error making http request: ", err)
			return
		}

		//response
		b, readErr := io.ReadAll(res.Body)
		if readErr != nil {
			fmt.Println(readErr)
			return
		}
		defer res.Body.Close()

		var resp PurchaseBody
		json.Unmarshal([]byte(b), &resp)

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "totalPrice": resp.TotalPrice})
	}
}
