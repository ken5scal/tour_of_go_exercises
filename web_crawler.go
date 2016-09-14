package main

import "fmt"

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls[]string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a max of depth
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel/
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either"
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

func main() {
	Crawl("http://golang.org/", 4, fetcher)
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
	return "", nil, fmt.Errorf("not found: %s",url)
}
