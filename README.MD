## Description:

Full-stack application using
- Golang/Gin
- React/Typescript
- MongoDB

---

## Server

### Set Up

Create a .env file in server directory and add these values
- `PORT`
- `MONGODB_URL`

### Development

- `go mod tidy`
- `go build main.go`
- `go run main.go`

---

## Client

### Development

- `yarn`
- `yarn dev`

---

## Status

- Server: MVP finished
- Client: MVP finished

---

## Changelog

### [09-05-2022]

- start up new server directory
  - `go.mod`
  - `main.go` creates and starts a gin server with route paths
  - `models` folder
    - `order.go`
  - `routes` folder
    - `connections.go` creates connection with MongoDB client
    - `orders.go` defines all the functions for each route

### [09-06-22]

- start up new client directory with `yarn create vite`
  - installed axios and react-bootstrap
  - `components` dir
    - `orders`
    - `single-order`
