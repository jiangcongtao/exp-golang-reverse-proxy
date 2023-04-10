# Build and Run

## Build

```shell
go mod tidy
go build
```

## Run

```shell
./exp-golang-reservse-proxy
```

## Test

In another terminal

```json
(base) ➜  ~ curl http://localhost:8080/get -v
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /get HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.86.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 502 Bad Gateway
< Date: Mon, 10 Apr 2023 03:44:42 GMT
< Content-Length: 0
<
* Connection #0 to host localhost left intact
(base) ➜  ~ curl http://localhost:8080/get -v
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> GET /get HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.86.0
> Accept: */*
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Access-Control-Allow-Credentials: true
< Access-Control-Allow-Origin: *
< Content-Length: 349
< Content-Type: application/json
< Date: Mon, 10 Apr 2023 03:45:10 GMT
< Server: gunicorn/19.9.0
<
{
  "args": {},
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Host": "localhost",
    "User-Agent": "curl/7.86.0",
    "X-Amzn-Trace-Id": "Root=1-64338646-340b5c181972b72e02a79655",
    "X-Forwarded-Host": "",
    "X-Origin-Host": "httpbin.org"
  },
  "origin": "127.0.0.1, 203.79.37.5",
  "url": "https:///get"
}
* Connection #0 to host localhost left intact
(base) ➜  ~ curl -X POST http://localhost:8080/post -d'{"name":"nick", "email":"nick.jiang@hilton.com"}' -v
Note: Unnecessary use of -X or --request, POST is already inferred.
*   Trying 127.0.0.1:8080...
* Connected to localhost (127.0.0.1) port 8080 (#0)
> POST /post HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.86.0
> Accept: */*
> Content-Length: 48
> Content-Type: application/x-www-form-urlencoded
>
* Mark bundle as not supporting multiuse
< HTTP/1.1 200 OK
< Access-Control-Allow-Credentials: true
< Access-Control-Allow-Origin: *
< Content-Length: 570
< Content-Type: application/json
< Date: Mon, 10 Apr 2023 03:45:17 GMT
< Server: gunicorn/19.9.0
<
{
  "args": {},
  "data": "",
  "files": {},
  "form": {
    "{\"name\":\"nick\", \"email\":\"nick.jiang@hilton.com\"}": ""
  },
  "headers": {
    "Accept": "*/*",
    "Accept-Encoding": "gzip",
    "Content-Length": "48",
    "Content-Type": "application/x-www-form-urlencoded",
    "Host": "localhost",
    "User-Agent": "curl/7.86.0",
    "X-Amzn-Trace-Id": "Root=1-6433864d-4033748178d610af5b09fbd5",
    "X-Forwarded-Host": "",
    "X-Origin-Host": "httpbin.org"
  },
  "json": null,
  "origin": "127.0.0.1, 203.79.37.5",
  "url": "https:///post"
}
* Connection #0 to host localhost left intact
```
