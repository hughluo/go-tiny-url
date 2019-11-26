# api

RESTful API for `tinyURL`

## Features

### Retreive Original URL

GET /gotinyurl/`tinyurl`

### Create Tiny URL

POST /gotinyurl/

FORM: "longURL": "`longURL`"

## Environment Variable Examples

### API_ADDRESS

":8080"

### GRPC_DIAL_TARGET

"localhost:3000"

### REDIS_MAIN_ADDRESS

"localhost:7001"

### REDIS_MAIN_PASSWORD

"supersecretpassword"
