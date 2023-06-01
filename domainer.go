package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func readTLDsFromFile(filePath string) ([]string, error) {
	// Open the file
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	// Read file contents
	content, err := io.ReadAll(file)
	if err != nil {
		return nil, err
	}

	// Split contents into lines
	lines := strings.Split(string(content), "\n")

	// Extract TLDs from lines
	var tlds []string
	for _, line := range lines {
		tld := strings.TrimSpace(line)
		tld = strings.TrimPrefix(tld, ".")
		if tld != "" {
			tlds = append(tlds, tld)
		}
	}

	return tlds, nil
}

func ExtractRootDomain(domain string, tlds []string) string {
	// Check if the domain ends with one of the passed TLDs
	for _, tld := range tlds {
		if strings.HasSuffix(domain, "."+tld) {
			// Remove the TLD from the domain
			domainWithoutTLD := strings.TrimSuffix(domain, "."+tld)
			domainWithoutTLD = removeScheme(domainWithoutTLD)
			// Split the domain into labels
			labels := strings.Split(domainWithoutTLD, ".")
			// Validate each label
			for _, label := range labels {
				if !isValidLabel(label) {
					return "none"
				}
			}
			// Extract the last label
			lastLabel := labels[len(labels)-1]
			return lastLabel + "." + tld
		}
	}

	return "none"
}

func removeScheme(domain string) string {
	schemeEndIndex := strings.Index(domain, "://")
	if schemeEndIndex != -1 {
		return domain[schemeEndIndex+3:]
	}
	return domain
}

func isValidLabel(label string) bool {
	// RFC implementation of valid domain name labels
	// See: https://cs.uwaterloo.ca/twiki/view/CF/HostNamingRules

	// Check the length of the label
	if len(label) > 63 {
		return false
	}
	// Check if the label starts with a letter or digit (RFC1123)
	if !unicode.IsLetter(rune(label[0])) && !unicode.IsDigit(rune(label[0])) {
		return false
	}
	// Check if the label ends with a letter or digit
	if !unicode.IsLetter(rune(label[len(label)-1])) && !unicode.IsDigit(rune(label[len(label)-1])) {
		return false
	}
	// Check if the interior characters of the label are letters, digits, or hyphen
	for _, char := range label[1 : len(label)-1] {
		if !unicode.IsLetter(char) && !unicode.IsDigit(char) && char != '-' {
			return false
		}
	}
	return true
}

func ReadTldsFromFile() []string {
	filePath := "tlds.txt"
	tlds, err := readTLDsFromFile(filePath)
	if err != nil {
		fmt.Println("Error reading TLDs:", err)
		os.Exit(1)
	}

	return tlds
}

func main() {
	stat, _ := os.Stdin.Stat()
	if (stat.Mode() & os.ModeCharDevice) != 0 {
		fmt.Fprintln(os.Stderr, "Please provide domains via stdin.")
		os.Exit(1)
	}

	tlds := ReadTldsFromFile()

	scanner := bufio.NewScanner(os.Stdin)
	seenDomains := make(map[string]bool)

	for scanner.Scan() {
		domain := scanner.Text()
		rootDomain := ExtractRootDomain(domain, tlds)
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
