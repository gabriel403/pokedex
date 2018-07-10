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

```
docker build .
```

```
Step 17/23 : CMD ["./pokedex-parser"]
 ---> Running in 85a651ce61a9
Removing intermediate container 85a651ce61a9
 ---> a389f02bee88
docker tag a389f02bee88 gabriel403/pokedex-parser:latest

Step 23/23 : CMD ["./pokedex-api"]
 ---> Running in 23a2e95c891e
Removing intermediate container 23a2e95c891e
 ---> 85409a02b3bc
docker tag 85409a02b3bc gabriel403/pokedex-api:latest

docker tag gabriel403/pokedex-parser:latest dalila.minders.us:5000/gabriel403/pokedex-parser:latest
docker tag gabriel403/pokedex-api:latest dalila.minders.us:5000/gabriel403/pokedex-api:latest
```

```
docker run --name pokedex-api -d --restart=always -it -p 5001:8080 -e POKEDEX_HTTP_TLS_CERTIFICATE=/certs/golandcertfile.crt -e POKEDEX_HTTP_TLS_KEY=/certs/server.key -v /home/gabriel/minders.us:/certs -v /home/gabriel/pokedex-assets:/go/src/pokedex/assets gabriel403/pokedex:latest
```
```
docker run -it --rm -e POKEDEX_ASSETS_DIR=/data/ -v /home/gabriel/pokedex-assets:/data gabriel403/pokedex-parser:latest
```
