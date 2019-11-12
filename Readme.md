# Simple CRUD App
# Requirements
- docker with multistage build capability
- optionally golang v1.9 or above for running test

# Usage
- start service: `docker-compose up -d --build`
- get single product: `curl http://0.0.0.0:3000/products/{sku}`
- get first 10 products: `curl http://0.0.0.0:3000/products`
- create product: `curl -X POST -d '{"sku":"foobar","name":"Some cool product","category":"A popular category"}' http://0.0.0.0:3000/products`
- update product: `curl -X PUT -d '{"sku":"foobar","name":"Some cool product","category":"A popular category"}' http://0.0.0.0:3000/products/foobar`
- delete product: `curl -X DELETE http://0.0.0.0:3000/products/{sku}`
