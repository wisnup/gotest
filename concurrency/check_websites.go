package concurrency

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {

	results := map[string]bool{}
	resultChan := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChan <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		u := <-resultChan
		results[u.string] = u.bool
	}

	return results
}
