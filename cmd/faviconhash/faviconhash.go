package main

import (
	"flag"
	"fmt"

	"github.com/binaryfigments/faviconhash"
)

func parseOptions() *faviconhash.Options {
	options := &faviconhash.Options{}

	flag.StringVar(&options.RequestURI, "uri", "https://hivemanager.databyte.nl/hm/images/favicon.ico", "Set Request URI")
	flag.StringVar(&options.UserAgent, "useragent", "Mozilla/5.0 (compatible; FaviconHash/0.1; +https://github.com/binaryfigments/faviconhash)", "Set User-Agent")
	flag.BoolVar(&options.Insecure, "insecure", true, "Allow insecure HTTPS requests")
	flag.Parse()

	return options
}

func main() {
	// Parse the command line flags and read config files
	options := parseOptions()

	hash, err := faviconhash.Run(options)
	if err != nil {
		println(err.Error())
	}
	fmt.Printf("Shodan Search Query: http.favicon.hash:%s\n", hash)
}
