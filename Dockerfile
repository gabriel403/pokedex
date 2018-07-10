FROM dalila.minders.us:5000/jgsqware/go-glide:latest as installer
COPY . /go/src/pokedex
WORKDIR /go/src/pokedex
RUN glide install

FROM dalila.minders.us:5000/golang:latest as api-builder
COPY --from=installer /go/src/pokedex /go/src/pokedex
WORKDIR /go/src/pokedex
RUN go build -o pokedex-api cmd/api/*.go

FROM dalila.minders.us:5000/golang:latest as parser-builder
COPY --from=installer /go/src/pokedex /go/src/pokedex
WORKDIR /go/src/pokedex
RUN go build -o pokedex-parser cmd/parser/*.go

FROM alpine:latest
COPY --from=parser-builder /go/src/pokedex/pokedex-parser /app/
VOLUME [ "/data" ]
WORKDIR /app/
CMD ["./pokedex-parser"]

FROM alpine:latest
COPY --from=api-builder /go/src/pokedex/pokedex-api /app/
EXPOSE 8080
VOLUME [ "/data" ]
WORKDIR /app/
CMD ["./pokedex-api"]
