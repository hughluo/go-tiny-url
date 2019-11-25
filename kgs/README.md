# kgs

key generation service for `tinyURL`

## Environment Variable Examples

### INIT_REDIS_FREE

Determine if init redis free, if not it will add all possible tinyURL in redis.

"false"

### KEY_LENGTH

The length of tinyURL

"6"

### REDIS_FREE_ADDRESS

"localhost:7001"

### REDIS_FREE_PASSWORD

"supersecretpassword"

### GRPC_LISTEN_PORT

":3000"
