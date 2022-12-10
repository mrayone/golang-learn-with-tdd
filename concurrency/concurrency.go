package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string // anonymous name
	bool
}

// race condition
/*
When a go routine performing write on object in memory in the same time.
To solve this problem we can use channels.
*/

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			// chan <- v send statement
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// <-chan receive expression
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
