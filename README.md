# faviconhash

[![GitHub license](https://img.shields.io/github/license/binaryfigments/faviconhash)](https://github.com/binaryfigments/faviconhash/blob/main/LICENSE.md)
[![Twitter](https://img.shields.io/twitter/url?style=social&url=binaryfigments)](https://twitter.com/binaryfigments)
[![GitHub Release](https://img.shields.io/github/release/binaryfigments/faviconhash)](https://github.com/binaryfigments/faviconhash/releases)

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

Something with versioning

```shell
git status
On branch main
Your branch is up to date with 'origin/main'.

Changes not staged for commit:
  (use "git add <file>..." to update what will be committed)
  (use "git restore <file>..." to discard changes in working directory)
        modified:   README.md

no changes added to commit (use "git add" and/or "git commit -a")

git add . && git commit -m "initial version"
[main 09a83d2] initial version
 1 file changed, 2 insertions(+), 1 deletion(-)

git tag -a v1.0.0 -m "initial version"

git push origin main --tags  
```

Go releaser test.