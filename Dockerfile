# syntax=docker/dockerfile:1
FROM golang:1.17-buster AS builder
WORKDIR /go/src/app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 GOOS=linux \
	 go build -o /go/bin/app ./cmd/

FROM gcr.io/distroless/static
COPY --from=builder /go/bin/app /go/bin/app
EXPOSE 8080
ENTRYPOINT ["/go/bin/app"]
