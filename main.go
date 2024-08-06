package main

import (
	"flag"
	"fmt"
	"golang.org/x/net/html"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

func scrape(baseURL string, frequencyMap *map[string]int, file *os.File) {
	res, err := http.Get(baseURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()

	if res.StatusCode > 399 {
		fmt.Fprintf(file, "%s Dead\n", baseURL)
		return
	}

	base, err := url.Parse(baseURL)
	if err != nil {
		log.Fatal(err)
	}

	z := html.NewTokenizer(res.Body)
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			if z.Err() == io.EOF {
				break
			}
			log.Fatal(z.Err())
		}
		if tt == html.StartTagToken {
			t := z.Token()
			if t.Data == "a" {
				for _, a := range t.Attr {
					if a.Key == "href" {
						href := a.Val
						link, err := url.Parse(href)

						if err != nil {
							log.Println(err)
							continue
						}
						absURL := base.ResolveReference(link).String()

						if (*frequencyMap)[absURL] == 0 {
							(*frequencyMap)[absURL] = 1
							if strings.Contains(absURL, "facebook.com") || strings.Contains(absURL, "twitter.com") || strings.Contains(absURL, "instagram.com") || strings.Contains(absURL, "youtube.com") {
								continue
							}
							scrape(absURL, frequencyMap, file)
						} else {
							(*frequencyMap)[absURL]++
						}
					}
				}
			}
		}
	}
}

func main() {

	file, err := os.Create("dead-links.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	urlPtr := flag.String("url", "", "The URL to scrape")
	flag.Parse()

	if *urlPtr == "" {
		log.Fatal("URL is required")
	}

	frequencyMap := make(map[string]int)
	scrape(*urlPtr, &frequencyMap, file)
}
