# Customer-Retailer Transaction APIs

This repository mimcs the transaction of products between customer and retailer

DataBase : mySql

Methods to start the repo:
* Pull the repo to the src folder of your go folder in home directory
* Start mySql server , and create database
* Update the Config file according to your database requirements and authentication
* start the project by the command in GoLand Terminal : go run main.go
* Check the status of the APIs using POSTMAN , and check the entries in the mySql database using mySql workbench

### APIs Documentation:

GET:  /product-api/products
> Get all the Products in the catalogue 

POST:  /product-api/product
> Create Product in the catalogue 

GET:  /product-api/product/:id
> Get Product by its ID from the catalogue 

UPDATE:  /product-api/product/:id
> Update Product by its ID from the catalogue

DELETE:  /product-api/product/:id
> Delete Product by its ID from the catalogue

POST:  /order-api/order
> order some product from the catalogue given product id and quantity in the POST request BODY

GET : /order-api/order/:id
> get the order request from the order id

GET :/order-api/orders
> get all Orders Information
