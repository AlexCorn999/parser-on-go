package actions

import (
	"fmt"
	"log"
	"net/http"
	"parser/telegram"

	"github.com/PuerkitoBio/goquery"
)

func HabrGo(tag string) {
	URL := fmt.Sprintf("https://habr.com/ru/hub/%s/top50/", tag)
	res, err := http.Get(URL)
	if err != nil {
		log.Fatalln(err)
	}

	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// parsing html
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	linkAll := doc.Find(".tm-article-snippet").Find(".tm-title_h2")
	link, _ := linkAll.Find("a").Attr("href")
	linkText, _ := linkAll.Find("span").Html()
	fmt.Println(linkText, link)

	//md5 hash create
	//linkMD5Sum := md5.Sum([]byte(link))

	text := fmt.Sprintf(`<b>Habr - %s</b>: <a href\=\"https://habr.com%s\">%s</a>`, tag, link, linkText)
	telegram.SendMessage(text)
	//fmt.Println(text)
	//db.CheckSiteNewBot(URL, link, text, fmt.Sprintf("%x", linkMD5Sum))
}
