# Hello Go
This repo creates a simple Golang server that responds to requests on localhost:8000/env with the message`Hello, ${WORLD}!`, where `${WORLD}` is the WORLD environment variable.

## Building the Dockerfile.
```bash
$ docker build -t hello-go .
```

## Running the image.
```bash
$ docker run -d -p 8000:8000 -e "WORLD=Go Programmer" --rm --name hello-go hello-go
```

## Executing a cURL command from within the container.
```bash
$ docker exec hello-go curl -s localhost:8000/env
Hello, Go Programmer! // Expected output
```

## Killing the container.
```bash
$ docker kill hello-go
```
