package main

import (
	"log"
	"net/http"
	"std/fmt"
	"strconv"

	// goquery - jquery처럼 css selector를 통해 원하는 요소를 쉽게 찾을 수 있게 도와줌
	"github.com/PuerkitoBio/goquery"
)

// 취업 정보 struct
type extractedJod struct {
	id       string
	title    string
	location string
	salay    string
	summary  string
}

// 취업 정보를 가져올 base URL
var baseURL string = "https://kr.indeed.com/jobs?q=python&limit=50"

func main() {
	totalPages := getPages()
	fmt.Println(totalPages)
	for i := 0; i < totalPages; i++ {
		getPage(i)
	}

}

// 페이지에 해당하는 URL을 return 하는 함수
func getPage(page int) {
	pageURL := baseURL + "&start=" + strconv.Itoa(page*50)
	fmt.Println("Requesting", pageURL)
	res, err := http.Get(pageURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	// 취업 정보를 가져옴; 취업 정보 카드
	searchCards := doc.Find(".jobsearch-SerpJobCard")

	searchCards.Each(func(i int, card *goquery.Selection) {
		// Attr - 데이터와 존재여부를 return
		id, _ := card.Attr("data-jk")
		fmt.Println(id)
		// Find - 원하는 속성을 가져옴
		title := card.Find(".title>a").Text()
		fmt.Println(title)
		location := card.Find(".sjcl").Text()
		fmt.Println(location)
	})
}

// 전체 페이지를 return 하는 함수
func getPages() int {
	pages := 0
	res, err := http.Get(baseURL)
	checkErr(err)
	checkCode(res)

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)
	checkErr(err)

	doc.Find(".pagination").Each(func(i int, s *goquery.Selection) {
		pages = s.Find("a").Length()
	})

	return pages
}

// 페이지를 가져오는데 Error가 발생하는지 체크
func checkErr(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}

// 페이지 코드를 체크
func checkCode(res *http.Response) {
	if res.StatusCode != 200 {
		log.Fatalln("Request failed with Status:", res.StatusCode)
	}
}
