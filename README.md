# go-tiny-url

Naive implementation of `tinyURL` service.

Four services included:

* `api`: API server to handle `tinyURL` request, gRPC client talk to kgs to fetch available `tinyURL`, redis client of `redis_main` to retrieve `longURL`(original URL)

* `kgs`: key generation service, gRPC server that distribute `tinyURL`, redis client of redis_free to get post generated `tinyURL`

* `redis_main`: storage key-value mapping of "`tinyURL`:`longURL`"

* `redis_free`: storage set of available `tinyURL`
