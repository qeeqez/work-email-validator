package workemailvalidator

import "golang.org/x/net/idna"

// domainToASCII converts any internationalized domain names to ASCII using Punycode.
// Reference: https://en.wikipedia.org/wiki/Punycode
func domainToASCII(domain string) string {
	asciiDomain, err := idna.ToASCII(domain)
	if err != nil {
		return domain
	}

	return asciiDomain
}
