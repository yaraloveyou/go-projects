package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"log"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

type SeoData struct {
	URL             string
	Title           string
	H1              string
	MetaDescription string
	StatusCode      int
}

type Parser interface {
	getSeoData(*http.Response) (SeoData, error)
}

type DefaultParser struct{}

var userAgents = []string{
	"Mozilla/5.0 (iPhone; CPU iPhone OS 11_9_3; like Mac OS X) AppleWebKit/603.20 (KHTML, like Gecko) Chrome/48.0.3572.116 Mobile Safari/601.4",
	"Mozilla/5.0 (compatible; MSIE 11.0; Windows; Windows NT 6.3; WOW64 Trident/7.0)",
	"Mozilla/5.0 (compatible; MSIE 10.0; Windows; Windows NT 6.3; en-US Trident/6.0)",
}

func main() {
	rand.Seed(time.Now().UnixNano()) // Инициализация генератора случайных чисел

	p := DefaultParser{}
	results := scrapeSiteMap("https://www.quicksprout.com/sitemap_index.xml", p, 10)
	for _, result := range results {
		fmt.Println(result)
	}
}

// Генерация случайного User-Agent
func randomUserAgent() string {
	randNum := rand.Intn(len(userAgents))
	return userAgents[randNum]
}

// Определение, является ли ссылка на карту сайта (XML) или страницу
func isSiteMap(urls []string) ([]string, []string) {
	sitemapFiles := []string{}
	pages := []string{}
	for _, page := range urls {
		if strings.Contains(page, "xml") {
			fmt.Println("Found sitemap:", page)
			sitemapFiles = append(sitemapFiles, page)
		} else {
			pages = append(pages, page)
		}
	}
	return sitemapFiles, pages
}

// Извлечение ссылок из карты сайта
func extractSiteMapURLs(url string) []string {
	worklist := make(chan []string, 1)
	toCrawl := []string{}
	var n int
	n++
	go func() {
		worklist <- []string{url}
	}()

	for n > 0 { // Исправлен цикл
		list := <-worklist
		n--
		for _, link := range list {
			n++
			go func(link string) {
				defer func() { n-- }() // Уменьшаем n при завершении горутины
				response, err := makeRequest(link)
				if err != nil {
					log.Printf("Error retrieving URL: %s", link)
					return
				}
				urls, err := extractUrls(response)
				if err != nil {
					log.Printf("Error extracting URLs from %s", link)
					return
				}
				sitemapFiles, pages := isSiteMap(urls)
				if len(sitemapFiles) > 0 {
					worklist <- sitemapFiles
				}
				toCrawl = append(toCrawl, pages...)
			}(link)
		}
	}

	return toCrawl
}

// Выполнение HTTP-запроса
func makeRequest(url string) (*http.Response, error) {
	client := http.Client{
		Timeout: time.Second * 10, // Увеличено время ожидания для стабильности
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", randomUserAgent())
	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	return res, nil
}

// Парсинг SEO-данных с использованием горутин и лимита конкурентных запросов
func scrapeURLs(urls []string, parser Parser, concurrency int) []SeoData {
	tokens := make(chan struct{}, concurrency)
	var n int
	worklist := make(chan []string, len(urls))
	results := []SeoData{}
	resultChan := make(chan SeoData, len(urls)) // Канал для сбора результатов

	go func() {
		worklist <- urls
	}()

	for n = len(urls); n > 0; n-- { // Исправлен цикл для правильного счета
		list := <-worklist
		for _, url := range list {
			if url != "" {
				go func(url string, tokens chan struct{}) {
					defer func() { n-- }()
					res, err := scrapePage(url, tokens, parser)
					if err != nil {
						log.Printf("Error scraping %s: %v", url, err)
					} else {
						resultChan <- res
					}
				}(url, tokens)
			}
		}
	}

	close(resultChan)
	for res := range resultChan {
		results = append(results, res)
	}
	return results
}

// Скрапинг одной страницы
func scrapePage(url string, token chan struct{}, parser Parser) (SeoData, error) {
	res, err := crawlPage(url, token)
	if err != nil {
		return SeoData{}, err
	}
	data, err := parser.getSeoData(res)
	if err != nil {
		return SeoData{}, err
	}
	return data, nil
}

// Запрос страницы с контролем количества горутин через токены
func crawlPage(url string, tokens chan struct{}) (*http.Response, error) {
	tokens <- struct{}{}
	defer func() { <-tokens }()
	resp, err := makeRequest(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

// Извлечение URL из XML с помощью goquery
func extractUrls(response *http.Response) ([]string, error) {
	doc, err := goquery.NewDocumentFromResponse(response)
	if err != nil {
		return nil, err
	}
	var urls []string
	doc.Find("loc").Each(func(i int, s *goquery.Selection) {
		url := s.Text()
		urls = append(urls, url)
	})
	return urls, nil
}

// Реализация парсера для извлечения SEO-данных
func (d DefaultParser) getSeoData(resp *http.Response) (SeoData, error) {
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		return SeoData{}, err
	}
	result := SeoData{
		URL:             resp.Request.URL.String(),
		StatusCode:      resp.StatusCode,
		Title:           doc.Find("title").First().Text(),
		H1:              doc.Find("h1").First().Text(),
		MetaDescription: doc.Find("meta[name^=description]").AttrOr("content", ""),
	}
	return result, nil
}

// Основная функция для скрапинга карты сайта
func scrapeSiteMap(url string, parser Parser, concurrency int) []SeoData {
	results := extractSiteMapURLs(url)
	return scrapeURLs(results, parser, concurrency)
}
