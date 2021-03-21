# 4 russians algorithm implemenation

## Content
- [Introduction](#Introduction)
- [How to run](#How-to-run)
  - [Prerequisites](#Prerequisites)
  - [Run](#Run)
- [Introduction](#Introduction)
- [Example requests](#Example-requests)

## Introduction

This repository contains implementation of 4 russians algorithm.
Description and additional information on this algorithm can be found on [Wikipedia](https://en.wikipedia.org/wiki/Method_of_Four_Russians#:~:text=In%20computer%20science%2C%20the%20Method,bounded%20number%20of%20possible%20values)

Pseudo algorithm can be found [here](https://louridas.github.io/rwa/assignments/four-russians/)

## How to run

### Prerequisites

- golang v1.16 or higher

### Run

1. go run main.go
2. send this POST [requests](#Example-requests) using curl or any other way

## Example requests
```bash
curl --location --request POST 'localhost:8080/matrix' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Left": [
    [false, true, false],
    [true, false, true],
    [false, true, false]
  ],
  "Right": [
    [true, false, true],
    [false, true, false],
    [true, false, true]
  ]
}'
```

```
curl --location --request POST 'localhost:8080/matrix' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Left": [
    [false, false, false],
    [false, false, false],
    [true, true, true]
  ],
  "Right": [
    [false, false, true],
    [true, false, true],
    [false, false, true]
  ]
}'
```

```
curl --location --request POST 'localhost:8080/matrix' \
--header 'Content-Type: application/json' \
--data-raw '{
  "Left": [
    [false, false, true],
    [true, false, true],
    [false, false, true]
  ],
  "Right": [
    [false, false, false],
    [false, false, false],
    [true, true, true]
  ]
}'
```