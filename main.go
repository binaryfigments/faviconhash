package faviconhash

import (
	"bytes"
	"crypto/tls"
	"encoding/base64"
	"fmt"
	"hash"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/twmb/murmur3"
)

// Options struct for running options
type Options struct {
	RequestURI string
	UserAgent  string
	Insecure   bool
}

// Run function that gets the hash
func Run(options *Options) (string, error) {
	favicon, err := faviconFromURL(options.RequestURI, options.Insecure, options.UserAgent)
	if err != nil {
		return "", err
	}
	return mmh3Hash32(standBase64(favicon)), nil
}

func faviconFromURL(requrl string, insecure bool, useragent string) (content []byte, err error) {
	_, err = url.ParseRequestURI(requrl)
	if err != nil {
		return nil, err
	}

	client := &http.Client{
		Timeout: time.Second * time.Duration(10),
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: insecure},
		},
	}

	req, err := http.NewRequest("GET", requrl, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", useragent)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	body, err := ioutil.ReadAll(resp.Body)

	resp.Body.Close()
	if err != nil {
		return nil, err
	}
	return body, nil
}

func mmh3Hash32(raw []byte) string {
	var h32 hash.Hash32 = murmur3.New32()
	h32.Write([]byte(raw))
	return fmt.Sprintf("%d", int32(h32.Sum32()))
}

func standBase64(braw []byte) []byte {
	bckd := base64.StdEncoding.EncodeToString(braw)
	var buffer bytes.Buffer
	for i := 0; i < len(bckd); i++ {
		ch := bckd[i]
		buffer.WriteByte(ch)
		if (i+1)%76 == 0 {
			buffer.WriteByte('\n')
		}
	}
	buffer.WriteByte('\n')
	return buffer.Bytes()
}
