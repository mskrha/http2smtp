[![Go Report Card](https://goreportcard.com/badge/github.com/mskrha/http2smtp)](https://goreportcard.com/report/github.com/mskrha/http2smtp)

## http2smtp

### Description
Simple HTTP to SMTP proxy.

### Build
```shell
git clone https://github.com/mskrha/http2smtp.git
cd http2smtp/source
make
```

### Build a Debian package
```shell
git clone https://github.com/mskrha/http2smtp.git
cd http2smtp/source
make deb
```

### Usage
```shell
curl \
	-D - \
	-X POST \
	http://127.0.0.1:8080/send/ \
	-d '{"from":"user1@local.domain","to":"user2@local.domain","subject":"Test","body":"This is a test message."}'
```
