package controllers

import (
	"cart/services"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
)

func Test() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func AddToCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product services.Product
		if bodyErr := ctx.ShouldBindBodyWith(&product, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := services.AddToCart(product); err != nil {
			fmt.Println("AddToCart: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetCartProducts() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var responseBody services.Product
		if bodyErr := ctx.ShouldBindBodyWith(&responseBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		message, products, err := services.GetCartProducts(responseBody)
		if err != nil {
			fmt.Println("GetCartProducts", err)
			ctx.JSON(http.StatusInternalServerError, gin.H{})
			return
		}

		ctx.JSON(200, gin.H{"message": message, "products": products})
	}
}

func ChangeProductQty() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product services.Product
		if bodyErr := ctx.ShouldBindBodyWith(&product, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := services.ChangeProductQty(product); err != nil {
			fmt.Println("ChangeProductQty: ", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func RemoveProductFromCart() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var productBody services.Product
		if bodyErr := ctx.ShouldBindBodyWith(&productBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if removeErr := services.RemoveProductFromCart(productBody); removeErr != nil {
			fmt.Println("RemoveProductFromCart: ", removeErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func ToPurchase() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var purchaseBody services.PurchaseBody
		if bodyErr := ctx.ShouldBindBodyWith(&purchaseBody, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		if err := services.AddTotalToCart(purchaseBody); err != nil {
			fmt.Println("AddTotalToCart", err)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK"})
	}
}

func GetTotalPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var response services.PurchaseBody
		if bodyErr := ctx.ShouldBindBodyWith(&response, binding.JSON); bodyErr != nil {
			fmt.Println("body: ", bodyErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		totalPrice, getErr := services.GetTotalPrice(response)
		if getErr != nil {
			fmt.Println("GetTotalPrice: ", getErr)
			ctx.Status(http.StatusInternalServerError)
			return
		}

		ctx.JSON(http.StatusOK, gin.H{"message": "OK", "totalPrice": totalPrice})
	}
}
