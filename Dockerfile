FROM golang:1.23-alpine

WORKDIR /app
COPY . .
RUN go mod tidy && go build -o key-value-cache
CMD ["./key-value-cache"]
EXPOSE 7171