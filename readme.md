```
docker run -it --rm -v $(pwd):/app -w /app jgsqware/go-glide glide install
```

```
docker run -it --rm -v $(pwd):/app -w /app jgsqware/go-glide glide get github.com/sirupsen/logrus
```

```
docker run -it --rm -v $(pwd):/go/src/pokedex -w /go/src/pokedex golang go run cmd/parser/pull-pokedex.go
```

```
docker run -it --rm -v $(pwd):/go/src/pokedex -p 8080:8080 -w /go/src/pokedex golang go run cmd/api/*.go
```
