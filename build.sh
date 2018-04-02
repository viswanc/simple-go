env GOOS=linux GOARCH=amd64 go build -o bin/simple-go src/main.go

docker build -t gcr.io/api-project-483163119575/simple-go .
