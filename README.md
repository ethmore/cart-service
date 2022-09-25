Cart Service
===

This GO project serves as a microservice for [eCommerce](https://github.com/users/ethmore/projects/4) project.


## Service tasks:

Create request to: 
- Retrieve cart info from [auth-and-db-service](https://github.com/ethmore/auth-and-db-service)
- Add product to cart database
- Remove product from cart databas
- Change product quantity
- Invoke purchase process



# Installation

Ensure GO is installed on your system
```
go mod download
````

```
go run .
```

## Test
```
curl http://localhost:3002/test
```
### It should return:
```
StatusCode        : 200
StatusDescription : OK
Content           : {"message":"OK"}
```

## Example .env file
This file should be placed inside `dotEnv` folder
```
# Cors URLs
BFF_URL = http://localhost:3001

# Request URLs
ADD_PRODUCT_TO_CART = http://127.0.0.1:3002/addProductToCart
ADD_TOTAL_TO_CART = http://127.0.0.1:3002/addTotalToCart
CHANGE_PRODUCT_QTY = http://127.0.0.1:3002/changeProductQty
GET_CART_PRODUCTS = http://127.0.0.1:3002/getCartProducts
GET_PRODUCT = http://127.0.0.1:3002/getProduct
GET_TOTAL_PRICE = http://127.0.0.1:3002/getTotalPrice
REMOVE_PRODUCT_FROM_CART = http://127.0.0.1:3002/removeProductFromCart
```