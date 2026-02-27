package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// As this function is a closure, the crawl function can access the local variables of the Crawl function.
	// Mutex and WaitGROup to synchronize access to fetched.
	var mu sync.Mutex
	var wg sync.WaitGroup

	fetched := make(map[string]bool)

	var crawl func(string, int)
	crawl = func(url string, depth int) {
		defer wg.Done()

		// Max depth reached.
		if depth <= 0 {
			return
		}

		// Lock Mutex and check if url already fetched.
		mu.Lock()
		if fetched[url] {
			// URL already fetched, release lock and return.
			mu.Unlock()
			return
		}
		// Remember URL as fetched and release lock.
		fetched[url] = true
		mu.Unlock()

		// Fetch content from URL.
		body, urls, err := fetcher.Fetch(url)
		if err != nil {
			// Could not fetch, so set URL to false in fetched, another worker might try again.
			mu.Lock()
			fetched[url] = false
			mu.Unlock()
			fmt.Println(err)
			return
		}
		fmt.Printf("found: %s %q\n", url, body)

		// Recursively crawl URLs one level deeper.
		for _, u := range urls {
			wg.Add(1)
			go crawl(u, depth-1)
		}
	}

	wg.Add(1)
	go crawl(url, depth)
	wg.Wait()
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func CrawlOld(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
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

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
