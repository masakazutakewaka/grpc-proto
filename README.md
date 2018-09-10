# What's this?

A sample implementation of micro services written in Go with Protocal Buffer, gRPC and grpc-gateway.

`#TODO: put on an architecture image`

# Execute following commands to build up services
```
$ dep init
$ docker-compose up --build client
```

open `localhost:8000` in your browser.

# Endpoints
## Items
- `GET: /item/{id}`
- `POST: /item`

## Users
- `GET: /user/{id}`
- `POST: /user`

## Coordinates
- `GET: /item/{itemId}/coordinates`
- `POST: /coordinate`

