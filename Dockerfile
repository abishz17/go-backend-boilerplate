FROM golang:1.22.2-alpine3.18

WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download && go mod verify

COPY . .

RUN apk update
RUN apk add git
RUN apk add curl
RUN go install -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate@latest

RUN cd / && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s
RUN rm -rf /var/cache/apk/*

CMD ["air", "-c", ".air.toml"]
