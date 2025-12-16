package workemailvalidator

import (
	_ "embed"
	"strings"
)

//go:embed data/disposable_domains.txt
var disposableDomainsData string

//go:embed data/free_domains.txt
var freeDomainsData string

var (
	disposableDomains map[string]bool
	freeDomains       map[string]bool
)

func init() {
	disposableDomains = loadDomains(disposableDomainsData)
	freeDomains = loadDomains(freeDomainsData)
}

func loadDomains(data string) map[string]bool {
	domains := make(map[string]bool)
	lines := strings.Split(data, "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		domains[strings.ToLower(line)] = true
	}
	return domains
}

func normalizeDomain(domain string) string {
	return strings.ToLower(strings.TrimSpace(domain))
}

func IsDisposableDomain(domain string) bool {
	domain = normalizeDomain(domain)
	return disposableDomains[domain]
}

func IsFreeDomain(domain string) bool {
	domain = normalizeDomain(domain)
	return freeDomains[domain]
}

func IsDisposableOrFreeDomain(domain string) bool {
	return IsDisposableDomain(domain) || IsFreeDomain(domain)
}

func IsBusinessDomain(domain string) bool {
	return !IsDisposableOrFreeDomain(domain)
}
