# Store Service

<p align="center">
  <a href="https://golang.org/doc/go1.18">
    <img src="https://img.shields.io/badge/Go-1.18+-00ADD8?style=flat&logo=go">
  </a>
  <a href="https://echo.labstack.com/">
    <img src="https://img.shields.io/badge/Echo-v4+-00ADD8?style=flat">
  </a>
  <a href="#">
    <img src="https://img.shields.io/badge/MySQL-5.7+-75C46B?style=flat">
  </a>
</p>

Is a simple project golang application with Echo Go web framework using MySQL and JWT base authentication middleware


## Configured with
- [gorm](https://github.com/jinzhu/gorm): ORM library for Golang
- [jwt-go](https://github.com/dgrijalva/jwt-go): JSON Web Tokens (JWT) as middleware
- [godotenv](https://github.com/joho/godotenv): go dotenv library
- Go Modules
- Built-in **RequestID Middleware**
- Feature **MySQL 5.7**
- Environment support

### Installation

```
$ go get https://github.com/ifty123/simple_online_store.git
```

## Running Your Application

Rename .env.example to .env and place your database credentials and jwt secret key

```
$ mv .env.example .env
$ go run main.go
```

migrate and seed database
```
$ go run main.go -m=migrate -s=all
```
## Entities/Domain

- User
- Product
- Category
- Cart
- Transaction


## Documentation API
you can see postman documentation : https://www.postman.com/winter-meteor-307866/workspace/simple-store-service
or run the app, and see localhost:{port}/swagger/index.html