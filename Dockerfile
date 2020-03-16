FROM golang:1.14-alpine

LABEL maintainer="Rajeev Singh <rajeevhub@gmail.com>"

WORKDIR /app

COPY . /app

# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .
CMD ["./main"]