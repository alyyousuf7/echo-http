# Echo HTTP
A very minimal HTTP server that returns the requested method and path. That's it.

## Usage
```bash
$ go get github.com/alyyousuf7/echo-http
$ echo-http -port 8080
```
`-port`: listening port (default `80`)

Sample response:
```bash
$ curl "http://localhost:8080/hello-world?some-param=some-value"
GET /hello-world?some-param=some-value
```

---

Adding `Requested-Status` header to the request with the expected status code, will make the server to respond with
that status code. If not provided, it will always be `200`.

```bash
$ curl "http://localhost:8080/hello-world?some-param=some-value" -H "Requested-Status: 404" -v
*   Trying ::1...
* TCP_NODELAY set
* Connected to localhost (::1) port 80 (#0)
> GET /hello-world?some-param=some-value HTTP/1.1
> Host: localhost:8080
> User-Agent: curl/7.64.1
> Accept: */*
> Requested-Status: 404
>
< HTTP/1.1 404 Not Found
< Date: Mon, 31 May 2021 15:55:10 GMT
< Content-Length: 39
< Content-Type: text/plain; charset=utf-8
<
GET /hello-world?some-param=some-value
* Connection #0 to host localhost left intact
* Closing connection 0
```

## Why?
Recently, I have written some NGINX configurations and it uses lots of `proxy_pass` within the `location` blocks. I
wanted to test those configurations. Using `http-server` worked well when testing configurations manually, however I
wanted to write some automated test scripts that could do some sanity checks on basic paths. So this was both.
