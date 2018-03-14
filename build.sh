env GOOS=linux GOARCH=amd64 go build -ldflags "-X main.version=0.1" -o bin/simple-go src/main.go

docker build -t viswanathct/simple-go .
