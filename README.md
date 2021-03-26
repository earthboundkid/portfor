# Portfor <a href="https://pkg.go.dev/github.com/carlmjohnson/portfor"><img src="https://pkg.go.dev/badge/github.com/carlmjohnson/portfor.svg" alt="Go Reference"></a>
## The problem
You develop a lot of sites on your computer and can never remember if a particular site is on port 3000, 3001, 8000, 8080, etc. etc.

## The solution
Portfor returns a deterministic hash for a given argument suitable for use as an unpriviledged port number. As a special case, it trims `www.` or `local.` prefixes from its argument.

## Usage
```
$ portfor example.com
24250

$ portfor www.example.com
24250

$ portfor -h
portfor v0.21.1

portfor returns a deterministic hash for a given argument suitable for use as an
unpriviledged port number. As a special case, it trims "www." or "local." prefixes
from its argument.

Options:

  -l    output localhost address (change string format with $PORTFOR_LOCALHOST)

$ PORT=$(portfor example.com) run-dev-server.sh
# server starts up on port 24250

$ another-run-script.sh -p $(portfor example.com)
# server starts up on port 24250

$ open $(portfor -l example.com)
# web browser opens to http://localhost:24250
```

## Installation
Requires [Go](https://golang.org/) to be installed.

```
$ go install github.com/carlmjohnson/portfor@latest
```
