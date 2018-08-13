FROM golang:latest as build

COPY id_rsa /root/.ssh/
RUN ssh-keyscan github.com >> /root/.ssh/known_hosts
RUN git config --global url."git@github.com:".insteadOf "https://github.com/"

WORKDIR /go/src/github.com/buschky/devnull
COPY . .

RUN go get ./...
RUN rm -r /root/.ssh
RUN go build -ldflags "-linkmode external -extldflags -static" -a main.go

FROM alpine:latest
WORKDIR /app/
COPY --from=build /go/src/github.com/buschky/devnull/main .
CMD ["./main"]
