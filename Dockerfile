FROM golang:1.20.2

WORKDIR /app

# Install sqlboiler and the PostgreSQL driver
RUN go install github.com/volatiletech/sqlboiler/v4@latest
RUN go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@latest

ADD . .
##RUN go mod download

CMD ["go", "run", "main.go"]
