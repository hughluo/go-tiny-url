# go-tiny-url

Naive implementation of `tinyURL` service.

Four services included:

* `api`: API server to handle `tinyURL` request, gRPC client talk to kgs to fetch available `tinyURL`, redis client of `redis_main` to retrieve `longURL`(original URL)

* `kgs`: key generation service, gRPC server that distributes `tinyURL`, redis client of redis_free to get post-generated `tinyURL`

* `redis_main`: storage key-value mappings of "`tinyURL`:`longURL`"

* `redis_free`: storage set of available `tinyURL`

## REPO
[`api`](https://github.com/hughluo/go-tiny-url-api)
[`kgs`](https://github.com/hughluo/go-tiny-url-kgs)

## ISSUE
`kgs` need warm up to generate all available `tinyURL` before hosting.
