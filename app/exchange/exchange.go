package exchange

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const nbkzUrl = "https://nationalbank.kz/rss/rates_all.xml" //@todo : extract

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
	Item Item `json:"item"` //@todo  : xml+json ?
}

func main() { //@todo : remove
}

func parseXMLRate(data io.ReadCloser) ([]Item, error) {
	xmlData, err := ioutil.ReadAll(data) //@todo : deprecated
	if err != nil {
		log.Fatal(err)
	}

	var channel struct {
		Items []Item `xml:"channel>item"`
	}
	if err := xml.Unmarshal(xmlData, &channel); err != nil {
		log.Fatal(err)
	}
	return channel.Items, nil
}

func GetCurrentRate(cur string) (Rate, error) { //@todo
	// Parse XML data

	// Make an HTTP GET request
	resp, err := http.Get(nbkzUrl)
	if err != nil {
		log.Fatal(err) //@todo : return err
	}
	defer resp.Body.Close()

	// Read the response body
	items, err := parseXMLRate(resp.Body)
	fmt.Println()
	// Find and store the chosen item (AUD)
	var chosenItem Item //@todo

	for _, item := range items { // @todo: extract
		if item.Title == cur {
			chosenItem = item
			break
		}
	}
	result := Rate{
		Title:   chosenItem.Title,
		Current: chosenItem.Description,
	}
	return result, nil

}
