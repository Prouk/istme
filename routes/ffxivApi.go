package routes

import (
	"encoding/json"
	"github.com/PuerkitoBio/goquery"
	"log"
	"net/http"
)

type FFXIVChar struct {
	CharUrl      string
	ImgUrl       string
	Title        string
	JobImg       string
	Level        string
	GrandCompany string
}

func FfxivApi(wr http.ResponseWriter, req *http.Request) {
	req.ParseForm()
	q := req.Form.Get("name")
	worldName := req.Form.Get("world")
	searchUrl := "https://fr.finalfantasyxiv.com/lodestone/character/?q=" + q + "&worldname=" + worldName
	result, err := http.Get(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	charUrl, exist := FirstOfSearch(result)
	if !exist {
		println(charUrl)
	}
	searchUrl = "https://fr.finalfantasyxiv.com" + charUrl
	result, err = http.Get(searchUrl)
	if err != nil {
		log.Fatal(err)
	}
	charToReturn := GetCharInfos(result)
	charToReturn.CharUrl = searchUrl
	jsresp, err := json.Marshal(charToReturn)
	wr.Header().Set("Content-Type", "application/json")
	wr.Write(jsresp)
}

func FirstOfSearch(result *http.Response) (string, bool) {
	doc, err := goquery.NewDocumentFromReader(result.Body)
	if err != nil {
		log.Fatal(result)
	}
	return doc.Find(".entry__link").First().Attr("href")
}

func GetCharInfos(result *http.Response) FFXIVChar {
	doc, err := goquery.NewDocumentFromReader(result.Body)
	if err != nil {
		log.Fatal(err)
	}
	char := FFXIVChar{}
	img, _ := doc.Find(".character__detail__image").Find("img").First().Attr("src")
	title := doc.Find(".frame__chara__title").Text()
	jobImg, _ := doc.Find(".character__class_icon").Find("img").First().Attr("src")
	level := doc.Find(".character__class__data").Find("p").First().Text()
	grandCompany := doc.Find(".character__freecompany__name").Find("a").First().Text()
	char.ImgUrl = img
	char.Title = title
	char.JobImg = jobImg
	char.Level = level
	char.GrandCompany = grandCompany
	return char
}
