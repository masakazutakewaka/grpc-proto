# What's this?

A sample implementation of micro services written in Go with Protocal Buffer, gRPC and grpc-gateway.

![alt text](https://user-images.githubusercontent.com/26586593/45367854-24a68600-b61d-11e8-865a-43bfdd7bc431.png)

inspired by https://github.com/tinrab/spidey

# Execute following commands to build up services
```
$ dep init
$ docker-compose up --build gateway
```

open `localhost:8000` in your browser.

# HTTP endpoints
## Items
- `GET: /item/{id}`
- `POST: /item`
  - e.g. `$ curl -X POST localhost:8000/item -d '{"name": "glove", "price": 3000}'`

## Users
- `GET: /user/{id}`
- `POST: /user`
  - e.g. `$ curl -X POST localhost:8000/user -d '{"name": "conor"}'`

## Coordinates
- `GET: /item/{itemId}/coordinates`
- `POST: /coordinate`
  - e.g. `$ curl -X POST localhost:8000/coordinate -d '{"userId": 1, "itemIds": [2,3,4]}'`

