# Go TCP Proxy

A simple TCP proxy that forwards traffic from a local port to a remote address.

## Usage

```bash
go run main.go -l :8080 -r example.com:80
```

## Structure
- `main.go`: Entry point and listener loop.
- `proxy/handle.go`: Connection handling and data copying.
- `config/config.go`: Argument parsing.
