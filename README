# Simple fasthttp Program
* an attempt to re-produce 502 error code, related to issue https://github.com/valyala/fasthttp/issues/739
* Default port: 8010 (can be configured in Dockerfile)
* used sleep (> 2 mins) to simulate long running task, client gets 502 at 120s = 2 mins

1. Build docker image
2. Open terminal 1 and run container
3. fire up a request in another terminal window

## Build docker image
```bash
docker build -t fasthttp-server .
```
## Run container
```bash
docker run --rm --name server -p 8010:8010 fasthttp-server
```

## Fire up a request
```bash
curl http://localhost:8010/
```
## Monitor container logs
```bash
docker logs server -f
```

Clean up local env
```bash
docker image rm fasthttp-server
```