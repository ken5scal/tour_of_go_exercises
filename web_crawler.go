package main

import (
	"fmt"
	"sync"
	"errors"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls[]string, err error)
}

type Fetched struct {
	//url map[string] error
	url map[string]error
	mux sync.Mutex
}

func (f *Fetched) addValue(url string) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.url[url] = errors.New("url load in progress.")
}

func (f *Fetched) updateLoadingStatus(url string, e error) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.url[url] = e
}

func (f *Fetched) isAlreadyFetched(url string) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	_, ok := f.url[url]
	return ok
}

var fetched = Fetched{url: make(map[string]error)}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a max of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel without fetching the same URL twice.
	if depth <= 0 {
		return
	}

	if fetched.isAlreadyFetched(url) {
		//fmt.Printf("%s already fetched\n", url)
		return
	}

	fetched.addValue(url)

	_, urls, err := fetcher.Fetch(url)

	fetched.updateLoadingStatus(url, err)

	if err != nil {
		fmt.Println(err)
		return
	}

	//fmt.Printf("found: %s %q \n", url, body)
	done := make(chan bool)
	for i, u := range urls {
		fmt.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(urls), url, u)
		go func(url string) {
			Crawl(url, depth - 1, fetcher)
			done <- true
		}(u)
	}
	for i, u := range urls {
		fmt.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(urls), u)
		<-done
	}
	fmt.Printf("<- Done with %v\n", url)
}

//fkeFetcher is Fetcher that returns canned results
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {

	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is populated fakeFetcher
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}