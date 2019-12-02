package validator

import (
	"strings"
)

type EmailDomainAllowed struct {
	Value          string
	AllowedDomains []string
}

func (v EmailDomainAllowed) Validate() (bool, Message) {
	stringParts := strings.Split(v.Value, "@")
	if len(stringParts) < 2 {
		return false, Message("Email value doesnt contain '@'")
	}

	domain := stringParts[1]
	domain = strings.TrimSpace(domain)

	for _, allowedDomain := range v.AllowedDomains {
		if domain == allowedDomain {
			return true, ""
		}
	}

	return false, "Email domain not allowed"
}
