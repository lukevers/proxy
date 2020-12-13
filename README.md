# proxy

Network proxy sever, built on [goproxy](https://github.com/elazarl/goproxy).

## Allowlist

If no allowlist is found, all incoming connections will be allowed. The default allowlist is at `/etc/proxy.addresses.list`, but can be configured with the `-addresses` flag. The allowlist is a new line separated list of IP addresses.

## Building

Nothing special. Adding this here for myself for the correct architecture for my own memory for the raspberry pi I have running this on.

```bash
# Standard
go build

# Cross-compile for raspberry pi:
env GOOS=linux GOARCH=arm GOARM=5 go build
```