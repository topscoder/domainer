# Domainer - Extract Root Domains from a List of Domains

Domainer is a Go script that extracts the root domain from a list of domains based on a provided list of top-level domains (TLDs) that follows the rules of ARPANET host names as specified in RFC1034 and RFC1123. The list of Top-Level Domains is based on the [official TLD list](https://data.iana.org/TLD/tlds-alpha-by-domain.txt). Also it respects so called Country Code Top-Level Domains.

## Installation

To install Domainer, you can use the following command:

```shell
go install github.com/topscoder/domainer
```

This will install the Domainer script as an executable in your Go bin directory.

## Usage

To use Domainer, provide a list of domains via stdin (standard input). The script will extract the root domain for each domain and print the results to stdout (standard output).

### Example Usage:

```shell
cat domains.txt | domainer
```

The domains should be provided as separate lines in the `domains.txt` file.

## TLDs

Domainer uses an internal list of TLDs, so there is no need for the end user to provide a separate TLDs file. The internal list includes a comprehensive set of TLDs, including both generic TLDs (gTLDs) and country code TLDs (ccTLDs).

## Notes

- The script follows the rules for ARPANET host names as specified in RFC1034 and RFC1123.
- The list of Top-Level Domains is based on the [official TLD list](https://data.iana.org/TLD/tlds-alpha-by-domain.txt).
- It respects so called Country Code Top-Level Domains, like .co.uk and .com.mx
- Domainer removes the scheme (if present) from the input domain before extracting the root domain.
- If a domain doesn't match any of the internal TLDs or if the extracted root domain is determined to be invalid, it will be marked as "none".

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for more details.

## Credits

This script was created by [topscoder](https://github.com/topscoder) together with ChatGPT.