// package main

// import (
// 	"encoding/xml"
// 	"fmt"
// 	"io/ioutil"
// 	"log"
// 	"net/http"
// )

// type RSS struct {
// 	XMLName xml.Name `xml:"rss"`
// 	Channel Channel  `xml:"channel"`
// }

// type Channel struct {
// 	Generator   string `xml:"generator"`
// 	Title       string `xml:"title"`
// 	Link        string `xml:"link"`
// 	Description string `xml:"description"`
// 	Language    string `xml:"language"`
// 	Copyright   string `xml:"copyright"`
// 	Items       []Item `xml:"item"`
// }

// type Item struct {
// 	Title       string `xml:"title"`
// 	PubDate     string `xml:"pubDate"`
// 	Description string `xml:"description"`
// 	Quantity    int    `xml:"quant"`
// 	Index       string `xml:"index"`
// 	Change      string `xml:"change"`
// 	Link        string `xml:"link"`
// }

// func main() {
// 	url := "https://nationalbank.kz/rss/rates_all.xml"

// 	// Make an HTTP GET request
// 	resp, err := http.Get(url)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer resp.Body.Close()

// 	// Read the response body
// 	xmlData, err := ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Unmarshal the XML data
// 	var rss RSS
// 	err = xml.Unmarshal(xmlData, &rss)
// 	if err != nil {
// 		log.Fatal(err)
// 	}

// 	// Access the unmarshaled data
// 	fmt.Println("Title:", rss.Channel.Title)
// 	fmt.Println("Description:", rss.Channel.Description)
// 	fmt.Println("Items:")

// 	for _, item := range rss.Channel.Items {
// 		fmt.Println("  Title:", item.Title)
// 		fmt.Println("  PubDate:", item.PubDate)
// 		fmt.Println("  Description:", item.Description)
// 		fmt.Println("  Quantity:", item.Quantity)
// 		fmt.Println("  Index:", item.Index)
// 		fmt.Println("  Change:", item.Change)
// 		fmt.Println("  Link:", item.Link)
// 		fmt.Println()
// 	}
// }
