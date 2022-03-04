package main

import (
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

const defaultParallel = 10
const FlagParallel = "parallel"

func main() {
	parallel, urls := parseParams()
	printMD5HashResponses(urls, parallel)
}

func parseParams() (parallel int, urls []string) {
	flag.IntVar(&parallel, FlagParallel, defaultParallel, "the number of parallel requests.")
	flag.Parse()
	if parallel <= 0 {
		parallel = defaultParallel
	}
	urls = flag.Args()
	return parallel, urls
}

func printMD5HashResponses(urls []string, parallel int) {
	indexProcessing := len(urls) - 1
	if indexProcessing < 0 || parallel <= 0 {
		return
	}
	doneChan := make(chan struct{})
	var processing int
	for {
		if processing < parallel && indexProcessing >= 0 {
			processing++
			go func(url string) {
				hash, err := getMD5HashResponse(url)
				if err != nil {
					hash = err.Error()
				}
				fmt.Printf("%v %v\n", url, hash)
				doneChan <- struct{}{}
			}(urls[indexProcessing])
			indexProcessing--
			continue
		}

		<-doneChan
		processing--
		if indexProcessing < 0 && processing == 0 {
			return
		}
		continue
	}
}

func getMD5HashResponse(httpUrl string) (hash string, err error) {
	formattedUrl, err := url.Parse(httpUrl)
	if err != nil {
		return
	}
	if formattedUrl.Scheme == "" {
		formattedUrl.Scheme = "http"
	}

	resp, err := http.Get(formattedUrl.String())
	if err != nil {
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	hash = fmt.Sprintf("%x", md5.Sum(body))
	return
}
