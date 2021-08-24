#!/usr/bin/env bash
set -xe

# install all packages and dependencies
go get github.com/gin-gonic/gin
go get gopkg.in/go-playground/validator.v9
go get -u gorm.io/gorm
go get golang.org/x/crypto/bcrypt
go get github.com/joho/godotenv
go get github.com/golang-jwt/jwt/v4
go get -u github.com/aws/aws-sdk-go/...
go get github.com/gin-contrib/cors



#build command
go build -o bin/application main.go