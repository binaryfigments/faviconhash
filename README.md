# faviconhash

Get the favicon hash for fun and profit. (poc version)

Installation:

```shell
GO111MODULE=on go get -v github.com/binaryfigments/faviconhash/cmd/faviconhash
```

Help:

```shell
$ faviconhash -h
Usage of faviconhash:
  -insecure
    	Allow insecure HTTPS requests (default true)
  -uri string
    	Set Request URI (default "https://example.nl/favicon.ico")
  -useragent string
    	Set User-Agent (default "Mozilla/5.0 (compatible; FaviconHash/0.1; +https://github.com/binaryfigments/faviconhash)")
```