The goal of this project is to build a web scraper to detect dead links on a website. A dead link is defined as one that returns a status code in the range of 4xx or 5xx.
The web scraper is designed to recursively check every anchor tag on the website, traversing through all the pages that belong to the same domain. If a page belongs to a different domain, the page itself will be checked for validity, but none of the links on that page will be followed or checked.
Key Features:
Recursive Link Checking: The scraper recursively visits every page within the same domain, ensuring all internal links are checked.
Domain Filtering: Pages from different domains are checked for their status but their internal links are not followed, preventing unnecessary crawling.
Dead Link Detection: Links that return status codes in the range of 4xx or 5xx are identified and logged as dead links.
Single Check per Page: Each discovered page is only checked once to prevent infinite loops and redundant processing.
