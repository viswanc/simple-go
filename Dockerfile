FROM alpine:3.4

RUN apk -U add ca-certificates

EXPOSE 8080

COPY ./bin/simple-go /bin/simple-go

CMD ["simple-go"]
