FROM golang:1.23 as builder
RUN mkdir /app
WORKDIR /app
COPY go.mod go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags="-s -w" -trimpath -o fastique ./cmd/...

FROM alpine AS runner
RUN mkdir /app
RUN apk add --no-cache tzdata
COPY --from=builder /app/fastique /app
CMD ["/app/fastique"]
