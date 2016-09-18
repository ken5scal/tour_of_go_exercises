package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls[]string, err error)
}

type Fetched struct {
	//url map[string] error
	url map[string] bool
	mux sync.Mutex
}

func (f *Fetched) updateValue(url string) {
	f.mux.Lock()
	defer f.mux.Unlock()
	f.url[url] = true
}

func (f *Fetched) isAlreadyFetched(url string) bool {
	f.mux.Lock()
	defer f.mux.Unlock()
	if _, ok := f.url[url]; !ok {
		return false
	}
	return f.url[url]
}

func main() {
	Crawl("http://golang.org/", 4, fetcher)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a max of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel without fetching the same URL twice.
	urls := make(chan []string)

	if depth <= 0 {
		return
	}

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("found: %s %q \n", url, body)
	for _, u := range urls {
		Crawl(u, depth - 1, fetcher)
	}
	return
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
	//"http://golang.org/pkg/": &fakeResult{
	//	"Packages",
	//	[]string{
	//		"http://golang.org/",
	//		"http://golang.org/cmd/",
	//		"http://golang.org/pkg/fmt/",
	//		"http://golang.org/pkg/os/",
	//	},
	//},
	//"http://golang.org/pkg/fmt/": &fakeResult{
	//	"Package fmt",
	//	[]string{
	//		"http://golang.org/",
	//		"http://golang.org/pkg/",
	//	},
	//},
	//"http://golang.org/pkg/os/": &fakeResult{
	//	"Package os",
	//	[]string{
	//		"http://golang.org/",
	//		"http://golang.org/pkg/",
	//	},
	//},
}