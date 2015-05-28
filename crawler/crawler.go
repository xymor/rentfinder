package main

import (
  "log"
  "fmt"
  "os"
  "strings"
  "github.com/PuerkitoBio/goquery"
  "github.com/olivere/elastic"
)

// Create a client and connect to http://192.168.2.10:9201
var client, err = elastic.NewClient(elastic.SetURL(os.Getenv("ES_URL")))


func RentScrape() {

  doc, errGet := goquery.NewDocument("http://rj.olx.com.br/rio-de-janeiro-e-regiao/zona-norte/imoveis/aluguel/apartamentos?q=tijuca+-barra&ret=1040") 
  if errGet != nil {
    log.Fatal(errGet)
  } 
  doc.Find(".section_OLXad-list ul.list li").FilterFunction(filter).Each(func(i int, s *goquery.Selection) {

    anchor := s.Find("a")

    rent := Rent{Id: anchor.AttrOr("name",""),
                Name: scrub(anchor.Text()),
                Price: scrub(s.Find(".OLXad-list-price").Text()),
                Link: anchor.AttrOr("href","")   ,
                Description: s.Find(".detail_specific").Text(),
                Images: []string{0:s.Find("img.image").AttrOr("src","")} ,
                Location: s.Find("i").Find(".detail_region").Text()}

    fmt.Printf("%v\n", rent)

    index(rent)

  })
}

func main() {
  RentScrape()
}

func scrub(s string) string{
  return strings.Join( strings.Fields(s), " ") 
}

func filter(i int, sel *goquery.Selection) bool{
    return true
    //return sel.Find(".col-4 p").First().Text() == "Hoje"
}

func index(rent Rent){
  _, err = client.Index().
      Index("rent").
      Type("rent").
      Id(rent.Id).
      BodyJson(rent).
      Do()
  if err != nil {
      // Handle error
      fmt.Printf("%v\n", rent)
      panic(err)
  }
}

type Rent struct {
    Id     string        `json:"id"`
    Name  string        `json:"name"`
    Link string           `json:"link"`
    Description    string        `json:"description,omitempty"`
    Location    string        `json:"location,omitempty"`
    Images    []string         `json:"images,omitempty"`
    Price    string        `json:"price,omitempty"`
}