FROM jgsqware/go-glide:latest as installer
COPY . /go/src/pokedex
WORKDIR /go/src/pokedex
RUN glide install

FROM golang:latest
COPY --from=installer /go/src/pokedex /go/src/pokedex
WORKDIR /go/src/pokedex
VOLUME ["/go/src/pokedex/assets"]
CMD ["go run cmd/api/*.go"]
