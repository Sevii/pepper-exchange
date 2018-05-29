#!/bin/bash


while [ true ]
do
curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "ask",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 30,
    "userId": "OTHERKID"
}'

curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "ask",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 35,
    "userId": "OTHERKID"
}'
sleep 1
curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "ask",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 40,
    "userId": "OTHERKID"
}'

curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "ask",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 45,
    "userId": "OTHERKID"
}'

curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "bid",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 30,
    "userId": "KID1"
}'
sleep 1
curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "bid",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 30,
    "userId": "KID1"
}'

curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "bid",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 30,
    "userId": "KID1"
}'

sleep 1

curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "bid",
    "exchange": "BTCLTC",
    "number": 777,
    "price": 35,
    "userId": "KID1"
}'
curl -X POST \
  http://localhost:8080/order \
  -H 'Cache-Control: no-cache' \
  -H 'Content-Type: application/json' \
  -H 'Postman-Token: 6a397958-f668-4240-b86e-66121ba3411e' \
  -d '{
    "direction": "bid",
    "exchange": "BTCLTC",
    "number": 7888,
    "price": 99,
    "userId": "KID1"
}'

done
