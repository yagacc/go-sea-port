FROM golang:1.15-alpine as builder
WORKDIR /build
COPY go.mod go.sum ./
COPY ./vendor ./vendor
COPY ./domain ./domain
COPY ./pkg ./pkg
ARG project
COPY $project .
RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -mod=vendor -a -o server .

FROM scratch
COPY --from=builder /build/server .
ENTRYPOINT ["./server"]
EXPOSE 3000 8000