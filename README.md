# go-socket-reuseport-example

This is a sample of an application built in Go using SO_REUSEPORT socket option.

![Go Version](https://img.shields.io/badge/go-1.19-blue.svg)

## How to Run

### Using [go-task](https://taskfile.dev/)

### Build

```sh
$ task build
task: [clean] rm -rf bin
task: [build] go build -o bin/server cmd/server/main.go
task: [build] go build -o bin/client cmd/client/main.go
```

### Server 1

Open a new terminal and execute the following command.

```sh
$ task server -- 1
```

### Server 2

Open a new terminal and execute the following command.

```sh
$ task server -- 2
```

### Check LISTEN Status

```sh
$ task ss
task: [ss] ss -atn | grep -F "0.0.0.0:9999"
LISTEN    0      128                0.0.0.0:9999                  0.0.0.0:*                                                                                     
LISTEN    0      128                0.0.0.0:9999                  0.0.0.0:*
```

### Client

Open a new terminal and execute the following command.

```sh
$ task client
task: [client] for i in {1..10}; do bin/client; done
RESPONSE FROM: 2
RESPONSE FROM: 1
RESPONSE FROM: 1
RESPONSE FROM: 1
RESPONSE FROM: 2
RESPONSE FROM: 1
RESPONSE FROM: 1
RESPONSE FROM: 2
RESPONSE FROM: 1
RESPONSE FROM: 1
```

### Stop servers

```sh
$ task close
```
