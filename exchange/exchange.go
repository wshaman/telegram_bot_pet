package exchange

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Rate struct {
	Title   string
	Current string
}

type Item struct {
	Title       string `xml:"title"`
	PubDate     string `xml:"pubDate"`
	Description string `xml:"description"`
	Quant       string `xml:"quant"`
	Index       string `xml:"index"`
	Change      string `xml:"change"`
}

type Response struct {
	Item Item `json:"item"`
}

func main() {
}

func Exchange(cur string) Rate {
	// Parse XML data

	url := "https://nationalbank.kz/rss/rates_all.xml"

	// Make an HTTP GET request
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Read the response body
	xmlData, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var channel struct {
		Items []Item `xml:"channel>item"`
	}
	if err := xml.Unmarshal([]byte(xmlData), &channel); err != nil {
		log.Fatal(err)
	}
	fmt.Println()
	// Find and store the chosen item (AUD)
	var chosenItem Item

	for _, item := range channel.Items {
		if item.Title == cur {
			chosenItem = item
			break
		}
	}
	result := Rate{
		Title:   chosenItem.Title,
		Current: chosenItem.Description,
	}
	return result

}
