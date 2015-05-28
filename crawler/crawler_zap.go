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

$('.content-minificha')

  doc, errGet := goquery.NewDocument("http://www.zapimoveis.com.br/aluguel/apartamentos/agr+rj++regiao-da-tijuca/#{%22precomaximo%22:%221600%22,%22parametrosautosuggest%22:[{%22Bairro%22:%22%22,%22Zona%22:%22%22,%22Agrupamento%22:%22Regi%C3%A3o%20da%20Tijuca%22,%22Estado%22:%22RJ%22}],%22pagina%22:1,%22paginaOrigem%22:%22ResultadoBusca%22,%22semente%22:%221832014754%22,%22formato%22:%22Lista%22,%22ordem%22:%22DataAtualizacao%22}") 
  if errGet != nil {
    log.Fatal(errGet)
  } 
  doc.Find("article#list div.box-default").FilterFunction(filter).Each(func(i int, s *goquery.Selection) {

    anchor := s.Find(".list-cell a")
    loc := []string{s.Find("span.bairro").Text(), s.Find("span.bairro").Text()}.Join(s, " - ")
    rent := Rent{Id: strings.Replace(s.Find("div.carousel").AttrOr("id","")), "prereleaseCarousel", "", -1),
                Name: scrub(anchor.Text()),
                Price: scrub(s.Find("span.price").Text()),
                Link: anchor.AttrOr("href","")   ,
                Description: loc,
                Images: s.Find("complex_slide img").AttrOr("src",""),
                Location: loc
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