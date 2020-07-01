package main

import (
	"fmt"
	"strings"

	"github.com/labstack/echo"
	"std/github.com/JunYeong-dev/Job-Scrapper/scrapper"
)

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
	term := strings.ToUpper(scrapper.CleanString(c.FormValue("term")))
	fmt.Println(term)
	return nil
}
