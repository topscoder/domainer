package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func extractRootDomain(domain string) string {
	parts := strings.Split(domain, ".")
	if len(parts) < 2 {
		return "none"
	}
	return strings.Join(parts[len(parts)-2:], ".")
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "Please provide domains via stdin.")
		os.Exit(1)
	}

	scanner := bufio.NewScanner(os.Stdin)
	seenDomains := make(map[string]bool)

	for scanner.Scan() {
		domain := scanner.Text()
		rootDomain := extractRootDomain(domain)
		if rootDomain != "none" {
			seenDomains[rootDomain] = true
		}
	}

	if scanner.Err() != nil {
		fmt.Fprintln(os.Stderr, "Error reading input:", scanner.Err())
		os.Exit(1)
	}

	for domain := range seenDomains {
		fmt.Println(domain)
	}
}

