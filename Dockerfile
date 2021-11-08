FROM golang:alpine AS build

RUN apk add --update git
WORKDIR /go/src/github.com/mstefa/arnold
COPY . .
RUN CGO_ENABLED=0 go build -o /go/bin/cmd/api/main.go

# Building image with the binary
FROM scratch
COPY --from=build /go/bin/codelytv-mooc-api /go/bin/arnold
ENTRYPOINT ["/go/bin/arnold"]