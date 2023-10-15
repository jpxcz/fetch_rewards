# Fetch Rewards
Interview process for fetch rewards

## Tests
Currently supporting the calculations for points. It's also on the Dockerfile.  
```
$ go test ./...
```

## Build and Run
```sh
$ docker build --tag rewards .
$ docker run  -p 127.0.0.1:8080:8080 rewards
```
