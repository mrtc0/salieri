# Amadeus Core (Go implementation)

Amadeus is Online Compiler on Hiroshima Institute of Technology.
This repository is an experimental Amadeus Go implementation.

---

# Setup

See [https://linuxcontainers.org/ja/lxd/getting-started-cli/](https://linuxcontainers.org/ja/lxd/getting-started-cli/)

```
sudo apt install lxd
newgrp lxd
sudo lxd init
lxc launch images:debian/jessie admin
lxc exec admin -- /bin/bash
apt update
apt install clang gcc build-essential
```

# Config

Edit `config/config.toml`.

```
[development]
bind = "127.0.0.1"
port = 8080
lxcname = "admin" # Container Name
```

# Launch Amadeus Core

On 127.0.0.1:8080
```
./server
```
or
```
go run server.go
```

# API

## POST /api/compiler/ 

- Code : Source Code
- Language : Compile Language
    - Current Support : clang
- Stdin : Stdin Text
- Stdout : Stdout
- Stderr : Stderr

### Example

```
curl -i -H 'Content-Type: application/json' -d @sample/simple_stdout.json localhost:8080/api/compiler/

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Sat, 14 Jan 2017 19:04:13 GMT
Content-Length: 204

{
  "Code": "#include \u003cstdio.h\u003e\nint main() {\n    printf(\"HELLO\\n\");\n    return 0;\n}\n",
  "Language": "clang",
  "Stdin": "",
  "Stdout": "HELLO\n",
  "Stderr": "0"
}
```

```
curl -i -H 'Content-Type: application/json' -d @sample/simple_stdout.json localhost:8080/api/compiler/

HTTP/1.1 200 OK
Content-Type: application/json; charset=utf-8
X-Powered-By: go-json-rest
Date: Sat, 14 Jan 2017 19:03:23 GMT
Content-Length: 235

{
  "Code": "#include \u003cstdio.h\u003e\nint main(){\nint i = 0;\nscanf(\"%d\", \u0026i);\nprintf(\"%d\\n\", i*2);\nreturn 0;\n}\n",
  "Language": "clang",
  "Stdin": "10\n",
  "Stdout": "20\n",
  "Stderr": "0"
}
```

# Test

```
go test -v ./
```
