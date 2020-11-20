# do-watch-cli

List resources from Digital Ocean

## Usage

```
do-watch-cli
```

## Download a latest binary

### Mac

```
curl -L https://github.com/ondrejsika/do-watch-cli-bin/raw/master/latest/do-watch-cli-darwin-amd64 -o do-watch-cli && chmod +x do-watch-cli && sudo mv do-watch-cli /usr/local/bin
```

### Linux

```
curl -L https://github.com/ondrejsika/do-watch-cli-bin/raw/master/latest/do-watch-cli-linux-amd64 -o do-watch-cli && chmod +x do-watch-cli && sudo mv do-watch-cli /usr/local/bin
```

## Build from source (go build)

```
make build
sudo mv do-watch-cli /usr/local/bin/do-watch-cli
```

## Changlog

- `v0.1.0` - First minimal working version - list droplets, volumes, load balancers, kubernetes
