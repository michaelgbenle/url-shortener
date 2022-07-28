package helpers

import "net/url"

func EnforceHTTP() string {
	if url[:4] !=  "http" {
		return "http://" + url
	}
	return url
}

func RemoveDomainError(url string) bool {

}