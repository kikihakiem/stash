# Simple CRUD App
# Requirements
- docker with multistage build capability
- optionally golang 1.9 for running test

# Usage
- start service: `docker-compose up -d --build`
- get single product: `curl http://0.0.0.0:3000/products/{sku}`
- get first 10 products: `curl http://0.0.0.0:3000/products`
- create product: `curl -d '{"sku":"foobar","name":"Some cool product","category":"A popular category"}' http://0.0.0.0:3000/products`
- update product: `curl -d '{"sku":"foobar","name":"Some cool product","category":"A popular category"}' http://0.0.0.0:3000/products/foobar`
- delete product: `curl http://0.0.0.0:3000/products/{sku}`
