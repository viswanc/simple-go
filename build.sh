env GOOS=linux GOARCH=amd64 go build -o bin/simple-go src/main.go

docker build -t viswanathct/simple-go .
