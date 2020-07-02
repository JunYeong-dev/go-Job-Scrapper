package main

import (
	"os"
	"strings"

	"github.com/labstack/echo"
	"std/github.com/JunYeong-dev/Job-Scrapper/scrapper"
)

const fileName string = "jobs.csv"

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":1323"))
}

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	// 프로젝트 파일 경로에 다운 받은 파일을 삭제함
	defer os.Remove(fileName)
	term := strings.ToUpper(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(term)
	// 파일 다운로드 - 앞: 파일 경로에 다운받은 파일의 파일명, 뒤: 웹에서 다운받아지는 파일의 파일명
	return c.Attachment(fileName, fileName)
}
