package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

type item struct {
	Name   string `json:"name"`
	Price  string `json:"price"`
	ImgUrl string `json:"imgurl"`
}

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("shopee.tw"),
	)

	c.OnHTML("div.stardust-tabs-panels", func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
		fmt.Println(h)
	})

	c.Visit("https://shopee.tw/")
}

//https://www.youtube.com/watch?v=NU4OlJVj1gs
