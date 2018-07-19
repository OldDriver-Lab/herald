
#build stage
FROM golang:alpine AS builder
WORKDIR /go/src/app
COPY . .
RUN echo "http://mirrors.ustc.edu.cn/alpine/v3.3/main/" > /etc/apk/repositories
RUN apk add --no-cache git
RUN go-wrapper download   # "go get -d -v ./..."
RUN go-wrapper install    # "go install -v ./..."

#final stage
FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /go/bin/app /app
ENTRYPOINT ./app
LABEL Name=herald Version=0.0.1
EXPOSE 6060
