package concurrency

type WebsiteChecker func(string) bool
type websiteCheckResult struct {
	websiteUrl  string
	websiteIsUp bool
}

func CheckWebsites(websiteChecker WebsiteChecker, websiteUrls []string) map[string]bool {
	results := make(map[string]bool)
	websiteCheckResultChannel := make(chan websiteCheckResult)

	for _, websiteUrl := range websiteUrls {
		go func(_websiteUrl string) {
			websiteCheckResultChannel <- websiteCheckResult{_websiteUrl, websiteChecker(_websiteUrl)}
		}(websiteUrl)
	}

	for i := 0; i < len(websiteUrls); i++ {
		result := <-websiteCheckResultChannel
		results[result.websiteUrl] = result.websiteIsUp
	}

	return results
}
