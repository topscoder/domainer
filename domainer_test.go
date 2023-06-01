package main

import (
	"testing"
)

func TestExtractRootDomain(t *testing.T) {
	tlds := ReadTldsFromFile()

	testCases := []struct {
		domain       string
		allowedTLDs  []string
		expectedRoot string
	}{
		{
			domain:       "foo.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "https://foo.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "ftp://foo.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "-foo.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "none",
		},
		{
			domain:       "1-1.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "11.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "1*1.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "none",
		},
		{
			domain:       "*.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "none",
		},
		{
			domain:       "foo-.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "none",
		},
		{
			domain:       "foo.example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "foo.example.foobar",
			allowedTLDs:  tlds,
			expectedRoot: "none",
		},
		{
			domain:       "example.com",
			allowedTLDs:  tlds,
			expectedRoot: "example.com",
		},
		{
			domain:       "example.co.uk",
			allowedTLDs:  tlds,
			expectedRoot: "example.co.uk",
		},
		{
			domain:       "example.de",
			allowedTLDs:  tlds,
			expectedRoot: "example.de",
		},
	}

	for _, testCase := range testCases {
		rootDomain := ExtractRootDomain(testCase.domain, testCase.allowedTLDs)
		if rootDomain != testCase.expectedRoot {
			t.Errorf("Domain %s -> expected %s, but got %s",
				testCase.domain, testCase.expectedRoot, rootDomain)
		}
	}
}
