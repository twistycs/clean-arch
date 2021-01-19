FROM golang:1.15-alpine



WORKDIR /app/clean-arch

COPY go.mod .
COPY go.sum .

RUN go mod tidy
COPY . .
RUN go build -o ./out/clean-arch .

EXPOSE 9999

CMD ["./out/clean-arch"]