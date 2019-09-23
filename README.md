## git service

### golang
depends on golang greater than 1.13

### dependencies
- gin-gonic/gin
- graphql-go/graphql
- joho/godotenv

### installation
``` bash
go get github.com/growerlab/letsgit-svc
cd $GOPATH/src/github.com/growerlab/letsgit-svc
go mod download
```

### build and run
`APP_ENV=dev go run main.go`

### env
- dev
- staging
- test
- production

### graphql
http://your-server-ip:port/graphql
