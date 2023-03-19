# Glowing-Potato

# Running Project

## Setup Database
- the database used is PostgreSQL latest version
- create database called glowing_potato
- go to root project directory
- open Makefile
- change database credential with your configuration
- run migration with `make up` command

## Run HTTP API
go to root project directory and run command below : 

```
$ gin -p 9091 -a 8081 serve-http

```

## Create Order Request
open postman, and create request to `localhost:9091/create-order` with post method and json payload, then copy payload below
```
{
    "customer_id": 1,
    "address_id": 3,
    "product_id": [1,2,3],
    "payment_method_id": 1
}
```
if error appears, please check the id sent in request payload.