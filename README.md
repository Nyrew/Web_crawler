
# Web Crawler

A concurrent web crawler implemented in Go, designed to traverse websites and report the number of internal links found for each page. 

## Features
- **Concurrency Control:** Limits the number of simultaneous requests to avoid overloading servers.
- **Page Limit:** Restricts the maximum number of pages crawled.
- **Domain Restriction:** Only crawls pages within the same domain as the base URL.
- **Detailed Report:** Generates a sorted report of pages and the number of internal links pointing to each.

---

## Project Structure

```
WEB_CRAWLER/
├── .gitignore
├── configure.go             # Configuration structure and utilities
├── crawl_page.go            # Core logic for crawling pages recursively
├── crawler/
│   └── (sub-packages if any)
├── get_html.go              # Fetches HTML content from URLs
├── get_urls_from_html.go    # Extracts URLs from HTML content
├── get_urls_from_html_test.go
├── go.mod                   # Go module dependencies
├── go.sum                   # Dependency checksums
├── main.go                  # Main entry point for the application
├── normalize_url.go         # Normalizes URLs for comparison
├── normalize_url_test.go
├── print_report.go          # Generates and prints sorted reports
├── print_report_test.go
├── README.md                # Project documentation (this file)
```

---

## Installation

Ensure you have Go installed on your system. If not, [install Go](https://go.dev/doc/install).

1. Clone the repository:

   ```bash
   git clone <repository-url>
   cd WEB_CRAWLER
   ```

2. Install dependencies:

   ```bash
   go mod tidy
   ```

3. Build the application:

   ```bash
   go build -o crawler
   ```

---

## Usage

Run the crawler with the following command:

```bash
./crawler URL maxConcurrency maxPages
```

- `URL`: The base URL to start crawling from (e.g., `https://example.com`).
- `maxConcurrency`: The maximum number of concurrent requests.
- `maxPages`: The maximum number of pages to crawl.

### Example

```bash
./crawler "https://example.com" 3 10
```

This will start crawling `https://example.com` with up to 3 concurrent requests and stop after crawling 10 pages.

---

## Features

1. **Multi-threaded Crawling**: Limits concurrency with a buffered channel.
2. **Configurable Crawling**: Set the maximum concurrency and pages via CLI arguments.
3. **HTML Parsing**: Extracts internal links from `<a>` tags and normalizes them.
4. **Report Generation**: Produces a sorted report showing pages with the number of internal links.

---

## Tests

The project includes unit tests for key components like URL normalization and URL extraction from HTML. To run the tests:

```bash
go test ./...
```

---

## Sample Output

During the crawl, each visited page is printed to the console. Once the crawl completes, a report is generated, for example:

```
=============================
  REPORT for https://example.com
=============================
Found 5 internal links to https://example.com/page1
Found 3 internal links to https://example.com/page2
Found 2 internal links to https://example.com/page3
```

---

## Contributing

Contributions are welcome! Please fork the repository and submit a pull request.

---

## License

This project is licensed under the MIT License. See the `LICENSE` file for details.
